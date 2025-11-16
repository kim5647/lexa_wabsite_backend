package service

import (
	"context"
	"errors"
	db "lexa_wabsite_backend/db/sqlc"
	"lexa_wabsite_backend/dto"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

type IUserRepository interface {
	Create(ctx context.Context, user db.User) (db.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (db.User, error)
}

type IAuthService interface {
	Register(ctx context.Context, input dto.RegisterRequest) (db.User, error)
	GenerateToken(userID int64, email string) (string, error)
	ValidateToken(tokenString string) (*UserClaims, error)
	Login(ctx context.Context, input dto.LoginRequest) (string, error)
}

type AuthService struct {
	UserRepository IUserRepository
	JWTConfig      JWTConfig
}

// NewAuthService - конструктор
func NewAuthService(repo IUserRepository) IAuthService {
	return &AuthService{
		UserRepository: repo,
		JWTConfig:      NewJWTConfig(),
	}
}

// GenerateToken - делегирует вызов статической функции GenerateToken
func (s *AuthService) GenerateToken(userID int64, email string) (string, error) {
	// Вызывает функцию, определенную, например, в jwt_logic.go
	return GenerateToken(userID, email, s.JWTConfig)
}

// ValidateToken - делегирует вызов статической функции ValidateToken
func (s *AuthService) ValidateToken(tokenString string) (*UserClaims, error) {
	// Вызывает функцию, определенную, например, в jwt_logic.go
	return ValidateToken(tokenString, s.JWTConfig)
}

// HashPassword - сама функция хэширования.
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

// ComparePassword - сравнивает предоставленный пароль с хэшем.
func (s *AuthService) ComparePassword(hashedPassword, password string) bool {
	// Возвращает nil, если пароли совпадают
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// RegisterNewUser - принимает DTO (данные от клиента) и создает пользователя.
// Я заменил 'user: repositiry.User' на DTO (обычно это происходит в Handler, но для примера)
func (s *AuthService) Register(ctx context.Context, input dto.RegisterRequest) (db.User, error) {
	// 1. Проверка бизнес-правил: существует ли пользователь?
	exists, err := s.UserRepository.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return db.User{}, err // Ошибка БД
	}
	if exists {
		return db.User{}, errors.New("пользователь с таким email уже зарегистрирован")
	}

	// 2. Хэшируем пароль
	hashedPassword, err := s.HashPassword(input.Password) // <-- ИСПОЛЬЗУЕМ input.Password!
	if err != nil {
		return db.User{}, err
	}

	// 3. Маппинг: Создаем модель БД (repository.User) с хэшем
	userModel := db.User{
		Name:         input.Name,
		Email:        input.Email,
		Phone:        input.Phone,
		HashPassword: hashedPassword,
	}

	// 4. Сохраняем через репозиторий
	createdUser, err := s.UserRepository.Create(ctx, userModel)
	if err != nil {
		return db.User{}, err
	}

	return createdUser, nil
}

// Login - проверяет учетные данные и генерирует JWT.
func (s *AuthService) Login(ctx context.Context, input dto.LoginRequest) (string, error) {
	// 1. Найти пользователя по Email
	user, err := s.UserRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		// Если пользователь не найден (например, ошибка pgx.ErrNoRows),
		// возвращаем общую ошибку для безопасности.
		return "", errors.New("неверный email или пароль")
	}

	// 2. Сравнить пароль
	if !s.ComparePassword(user.HashPassword, input.Password) {
		return "", errors.New("неверный email или пароль")
	}

	// 3. Сгенерировать токен (если аутентификация успешна)
	token, err := s.GenerateToken(int64(user.ID), user.Email)
	if err != nil {
		return "", errors.New("не удалось сгенерировать токен")
	}

	return token, nil
}
