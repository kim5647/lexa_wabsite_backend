package repository

// Импортируем сгенерированный SQLC-код, который вы назвали 'sqlc'
import sqlc "lexa_wabsite_backend/db/sqlc"

// UserRepository — это ваша пользовательская структура репозитория.
// Она должна содержать сгенерированные запросы.
type UserRepository struct {
	// Используем тип Queries из пакета sqlc
	sqlQueries *sqlc.Queries
}

// NewUserRepository — ЭКСПОРТИРУЕМЫЙ конструктор, который вызывает ваш main.go.
// Он принимает сгенерированный объект sqlc.Queries.
func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		sqlQueries: q,
	}
}
