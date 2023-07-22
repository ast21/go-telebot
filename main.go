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
type Course data.Course
type Question data.Question
type QuestionAnswer data.QuestionAnswer
type UserAnswer data.UserAnswer

var db, err = sqlx.Connect("postgres", env("DSN"))

func main() {
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
		courses := getAllCourses()

		for i := 0; i < len(courses); i++ {
			if c.Text() == courses[i].Title {
				question, err := getRandomQuestionByCourseId(courses[i].ID)
				if err != nil {
					return c.Send("Вопросы еще не созданы, попробуйте чуть позже")
				}
				return c.Send(question.Title)
			}
		}

		return c.Send(c.Text())
	})

	var (
		// Universal markup builders.
		menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
		selector = &tele.ReplyMarkup{}

		// Reply buttons.
		rows []tele.Row
		//btnHelp = menu.Text("ℹ Help")
		//btnSettings = menu.Text("⚙ Settings")

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

	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)

	b.Handle("/start", func(c tele.Context) error {
		user := getUserByTelegramId(c)

		question := Question{}
		sql := `                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            SELECT
            q.id            as "id",
            q.title         as "title",
            q.created_at    as "created_at",
            q.updated_at    as "updated_at",

            c.id            as "course.id",
            c.title         as "course.title",
            c.created_at    as "course.created_at",
            c.updated_at    as "course.updated_at"
            FROM questions as q JOIN courses as c ON q.course_id = c.id
            LIMIT 1
        `
		err = db.Get(&question, sql)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(question)
		fmt.Println(question.Title)
		fmt.Println(question.Course.Title)

		//fmt.Println(user)

		var ()

		courses := getAllCourses()
		rows = []tele.Row{}
		for i := 0; i < len(courses); i++ {
			rows = append(rows, menu.Row(menu.Text(courses[i].Title)))
		}

		menu.Reply(rows...)

		welcomeText := fmt.Sprintf("Hello %s %s! I am Learning Assistant. Welcome!", user.FirstName, user.LastName)
		return c.Send(welcomeText, menu)
	})

	// On reply button pressed (message)
	//b.Handle(&btnHelp, func(c tele.Context) error {
	//	return c.Edit("Here is some help: ...")
	//})

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

func printSlice(s []tele.Row) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func getUserByTelegramId(c tele.Context) User {
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

	return user
}

func getAllCourses() []Course {
	var courses []Course
	err = db.Select(&courses, `SELECT * FROM courses`)
	if err != nil {
		log.Fatalln(err)
	}

	return courses
}

func getAllQuestionByCourseId(courseId uint64) []Question {
	var questions []Question
	err = db.Select(&questions, `SELECT * FROM "questions" WHERE "question_id" = $1`, courseId)
	if err != nil {
		log.Fatalln(err)
	}

	return questions
}

func getRandomQuestionByCourseId(courseId uint64) (Question, error) {
	var question Question

	err = db.Get(&question, `
        SELECT * FROM "questions"
            WHERE "course_id" = $1
            ORDER BY RANDOM()
            LIMIT 1
    `, courseId)

	return question, err
}
