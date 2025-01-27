package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// InitDB инициализирует подключение к PostgreSQL.
func InitDB(connString string) error {
	var err error
	db, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	// Создаём таблицу tasks, если её нет
	query := `
        CREATE TABLE IF NOT EXISTS tasks (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            completed BOOLEAN NOT NULL DEFAULT FALSE,
            created_at TIMESTAMP NOT NULL DEFAULT NOW()
        );
    `
	_, err = db.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("не удалось создать таблицу tasks: %w", err)
	}

	return nil
}

func GetAllTasks() ([]Task, error) {
	rows, err := db.Query(context.Background(), "SELECT id, title, completed, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
