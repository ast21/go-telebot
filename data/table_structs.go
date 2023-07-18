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
