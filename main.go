package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
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
