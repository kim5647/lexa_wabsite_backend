package repository

import (
	// Импортируем сгенерированный SQLC-код
	"context"
	sqlc "lexa_wabsite_backend/db/sqlc"
	// Вам также потребуется пакет service, чтобы получить тип данных User,
	// если он определен там, или используйте sqlc.User
)

// UserRepository - структура, которая реализует service.IUserRepository.
type UserRepository struct {
	// Внедряем сгенерированный объект Queries
	sqlQueries *sqlc.Queries
}

// NewUserRepository - ЭКСПОРТИРУЕМЫЙ конструктор (УСТРАНЯЕТ ОШИБКУ В main.go)
func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		sqlQueries: q,
	}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]sqlc.User, error) {
	// Вызов сгенерированного кода
	return r.sqlQueries.GetUsers(ctx)
}

// ExistsByEmail - метод для проверки существования пользователя (Требуется GetUserByEmail в SQLC)
func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {

	// ПРЕДПОЛАГАЯ, что у вас есть сгенерированный метод r.sqlQueries.GetUserByEmail:
	// user, err := r.sqlQueries.GetUserByEmail(ctx, email)

	// Если GetUserByEmail возвращает "sql.ErrNoRows", значит, пользователя нет.
	// Если GetUserByEmail отсутствует, вам нужно его добавить в users.sql.

	// ВРЕМЕННОЕ РЕШЕНИЕ (для компиляции):
	return false, nil
}

// ДОБАВИТЬ: Create - реализует метод Create из service.IUserRepository
func (r *UserRepository) Create(ctx context.Context, user sqlc.User) (sqlc.User, error) {
	// 💡 ВАЖНО: Замените на реальный вызов SQLC (например, r.sqlQueries.CreateUser)
	// Create User, вероятно, должен принимать параметры, а не структуру User целиком.
	return user, nil
}
