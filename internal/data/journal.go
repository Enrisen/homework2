package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/cohune-cabbage/di/internal/validator"
)

type Journal struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Date      time.Time `json:"date"`
}

func ValidateJournal(v *validator.Validator, journal *Journal) {
	v.Check(validator.NotBlank(journal.Title), "title", "must be provided")
	v.Check(validator.MaxLength(journal.Title, 100), "title", "must not be more than 100 bytes long")
	v.Check(validator.NotBlank(journal.Content), "content", "must be provided")
	v.Check(validator.MaxLength(journal.Content, 5000), "content", "must not be more than 5000 bytes long")
	// Date validation is handled by the HTML date input
}

type JournalModel struct {
	DB *sql.DB
}

func (m *JournalModel) Insert(journal *Journal) error {
	query := `
		INSERT INTO journal (title, content, date)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(
		ctx,
		query,
		journal.Title,
		journal.Content,
		journal.Date,
	).Scan(&journal.ID, &journal.CreatedAt)
}
