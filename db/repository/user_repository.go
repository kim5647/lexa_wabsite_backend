package repository

import (
	// Импортируем сгенерированный SQLC-код
	"context"
	db "lexa_wabsite_backend/db/sqlc"

	"github.com/jackc/pgx/v5"
	// Вам также потребуется пакет service, чтобы получить тип данных User,
	// если он определен там, или используйте sqlc.User
)

// UserRepository - структура, которая реализует service.IUserRepository.
type UserRepository struct {
	// Внедряем сгенерированный объект Queries
	sqlQueries *db.Queries
}

// NewUserRepository - ЭКСПОРТИРУЕМЫЙ конструктор (УСТРАНЯЕТ ОШИБКУ В main.go)
func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{
		sqlQueries: q,
	}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]db.User, error) {
	// Вызов сгенерированного кода
	return r.sqlQueries.GetUsers(ctx)
}

// ExistsByEmail - Окончательная реализация
func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {

	// Вызов сгенерированного кода
	_, err := r.sqlQueries.GetUserByEmail(ctx, email)

	if err != nil && err == pgx.ErrNoRows {
		return false, nil // Пользователь не найден
	}
	if err != nil {
		return false, err // Ошибка БД
	}
	return true, nil // Пользователь найден
}

// Create - Окончательная реализация (УСТРАНЯЕТ ОШИБКУ 'missing method Create')
func (r *UserRepository) Create(ctx context.Context, user db.User) (db.User, error) {
	// 1. Создаем параметры из структуры User
	params := db.CreateUserParams{
		Name:         user.Name,
		HashPassword: user.HashPassword,
		Phone:        user.Phone,
		Email:        user.Email,
	}

	// 2. Вызываем сгенерированный SQLC-метод
	createdUser, err := r.sqlQueries.CreateUser(ctx, params)

	// 3. Возвращаем результат
	return createdUser, err
}
