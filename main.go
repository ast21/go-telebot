package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"telebot/data"
	"time"

	tele "gopkg.in/telebot.v3"

	_ "github.com/lib/pq"
)

type User data.User

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

	pref := tele.Settings{
		Token:  env("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello! How are you?")
	})

	// All the text messages that weren't
	// captured by existing handlers.
	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(c.Text())
	})

	var (
		// Universal markup builders.
		menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
		selector = &tele.ReplyMarkup{}

		// Reply buttons.
		btnHelp     = menu.Text("ℹ Help")
		btnSettings = menu.Text("⚙ Settings")

		// Inline buttons.
		//
		// Pressing it will cause the client to
		// send the bot a callback.
		//
		// Make sure Unique stays unique as per button kind
		// since it's required for callback routing to work.
		//
		btnPrev = selector.Data("Prev", "prev")
		btnNext = selector.Data("Next", "next")
	)

	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)
	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)

	b.Handle("/start", func(c tele.Context) error {
		sender := c.Sender()

		user := User{}
		err = db.Get(&user, "SELECT * FROM users WHERE telegram_id = $1", sender.ID)
		if err != nil {
			fmt.Println(err)

			tx := db.MustBegin()
			tx.MustExec("INSERT INTO users (telegram_id, first_name, last_name) VALUES ($1, $2, $3)", sender.ID, sender.FirstName, sender.LastName)
			err := tx.Commit()

			if err != nil {
				log.Fatalln(err)
			}

			err = db.Get(&user, "SELECT * FROM users WHERE telegram_id = $1", sender.ID)
			if err != nil {
				log.Fatalln(err)
			}
		}

		fmt.Println(user)

		welcome := fmt.Sprintf("Hello %s %s! I am Learning Assistant. Welcome!", user.FirstName, user.LastName)
		return c.Send(welcome, menu)
	})

	// On reply button pressed (message)
	b.Handle(&btnHelp, func(c tele.Context) error {
		return c.Edit("Here is some help: ...")
	})

	// On inline button pressed (callback)
	b.Handle(&btnPrev, func(c tele.Context) error {
		return c.Respond()
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
