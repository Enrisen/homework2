package data

import (
	"context"
	"database/sql"
	"time"
)

type Todo struct {
	ID        int64     `json:"id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	Version   int32     `json:"version"`
}

type TodoModel struct {
	DB *sql.DB
}

func (m TodoModel) Insert(todo *Todo) error {
	query := `
		INSERT INTO todos (task, completed)
		VALUES ($1, $2)
		RETURNING id, created_at, version`

	args := []interface{}{todo.Task, todo.Completed}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&todo.ID, &todo.CreatedAt, &todo.Version)
}

func (m TodoModel) GetAll() ([]*Todo, error) {
	query := `
		SELECT id, task, completed, created_at, version
		FROM todos
		ORDER BY created_at ASC`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*Todo{}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Task,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.Version,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (m TodoModel) Update(todo *Todo) error {
	query := `
		UPDATE todos
		SET task = $1, completed = $2, version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING version`

	args := []interface{}{
		todo.Task,
		todo.Completed,
		todo.ID,
		todo.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&todo.Version)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			// Handle concurrency conflict if needed
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m TodoModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM todos
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m TodoModel) Get(id int64) (*Todo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT id, task, completed, created_at, version
        FROM todos
        WHERE id = $1`

	var todo Todo

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&todo.ID,
		&todo.Task,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.Version,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &todo, nil
}
