package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"

	_ "github.com/lib/pq"
)

type User struct {
	ID         uint64    `db:"id"`
	TelegramId string    `db:"telegram_id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func main() {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", env("DSN"))
	if err != nil {
		log.Fatalln(err)
	}

	var users []User
	err = db.Select(&users, "SELECT * FROM users ORDER BY id ASC")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(users)

	pref := tele.Settings{
		Token:  env("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! I am Learning Assistant. Welcome!")
	})

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	// All the text messages that weren't
	// captured by existing handlers.
	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(c.Text())
	})

	b.Start()
}

func env(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
