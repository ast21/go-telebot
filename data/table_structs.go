package data

import "time"

type User struct {
	ID         uint64    `db:"id"`
	TelegramId string    `db:"telegram_id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type Course struct {
	ID        uint64    `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Question struct {
	ID        uint64    `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Course Course `db:"course"`
}

type QuestionAnswer struct {
	ID        uint64    `db:"id"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Question Question `db:"question"`
}

type UserAnswer struct {
	ID        uint64    `db:"id"`
	Value     string    `db:"value"`
	Correct   bool      `db:"correct"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	User     User     `db:"user"`
	Question Question `db:"question"`
}
