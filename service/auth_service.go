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
}

type IAuthService interface {
	// Добавьте методы, которые Handler будет вызывать
	Register(ctx context.Context, input dto.RegisterRequest) (db.User, error)
	// ...
}

// AuthService - структура для логики аутентификации.
type AuthService struct {
	// Добавляем ЗАВИСИМОСТЬ от репозитория
	UserRepository IUserRepository
}

// NewAuthService - конструктор
func NewAuthService(repo IUserRepository) IAuthService {
	return &AuthService{
		UserRepository: repo,
	}
}

// HashPassword - сама функция хэширования.
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
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
