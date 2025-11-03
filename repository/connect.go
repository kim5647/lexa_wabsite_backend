package db // Новый пакет для логики подключения к БД

import (
	"context"
	"fmt" // Для форматирования ошибок

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectPool устанавливает и возвращает пул соединений.
// Имя должно быть ConnectPool, чтобы его можно было вызвать в main.go.
func ConnectPool() (*pgxpool.Pool, error) {
	// В реальном проекте connStr берется из os.Getenv("DATABASE_URL")
	connStr := "postgres://postgres:3006@localhost:3006/lexa_group?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		// Возвращаем ошибку. Решение о log.Fatal принимает main.go.
		return nil, fmt.Errorf("ошибка создания пула соединений: %w", err)
	}

	// Пингуем, чтобы убедиться, что соединение установлено
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ошибка проверки соединения с БД: %w", err)
	}

	return pool, nil
}
