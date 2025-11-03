package service

import (
	"errors" // Для создания пользовательских ошибок (например, "пользователь уже существует")

	"golang.org/x/crypto/bcrypt"

	// Предполагаем, что вы создали пакет 'repository'
	"lexa_wabsite_backend/dto" // Предполагаем, что у вас есть DTO для регистрации
	"lexa_wabsite_backend/repository"
)

const bcryptCost = 12

// --- ИНТЕРФЕЙСЫ (Лучшая Практика) ---
// IUserRepository - определяет методы, которые AuthService использует в репозитории
type IUserRepository interface {
	Create(user repository.User) (repository.User, error)
	ExistsByEmail(email string) (bool, error)
}

// AuthService - структура для логики аутентификации.
type AuthService struct {
	// Добавляем ЗАВИСИМОСТЬ от репозитория
	UserRepository IUserRepository
}

// NewAuthService - конструктор
func NewAuthService(repo IUserRepository) *AuthService {
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
func (s *AuthService) RegisterNewUser(input dto.RegisterRequest) (repository.User, error) {
	// 1. Проверка бизнес-правил: существует ли пользователь?
	exists, err := s.UserRepository.ExistsByEmail(input.Email)
	if err != nil {
		return repository.User{}, err // Ошибка БД
	}
	if exists {
		return repository.User{}, errors.New("пользователь с таким email уже зарегистрирован")
	}

	// 2. Хэшируем пароль
	hashedPassword, err := s.HashPassword(input.Password) // <-- ИСПОЛЬЗУЕМ input.Password!
	if err != nil {
		return repository.User{}, err
	}

	// 3. Маппинг: Создаем модель БД (repository.User) с хэшем
	userModel := repository.User{
		Name:         input.Name,
		Email:        input.Email,
		Phone:        input.Phone,
		HashPassword: hashedPassword,
	}

	// 4. Сохраняем через репозиторий
	createdUser, err := s.UserRepository.Create(userModel)
	if err != nil {
		return repository.User{}, err
	}

	return createdUser, nil
}
