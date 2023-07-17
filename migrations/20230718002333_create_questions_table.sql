-- +goose Up
-- +goose StatementBegin
CREATE TABLE "questions" (
    "id" bigserial PRIMARY KEY,
    "course_id" bigserial,
    "title" varchar,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "question_answers" (
    "id" bigserial PRIMARY KEY,
    "question_id" bigserial,
    "value" varchar,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_answers" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigserial,
    "question_id" bigserial,
    "value" varchar,
    "correct" boolean,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);

ALTER TABLE "questions" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "question_answers" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

ALTER TABLE "user_answers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_answers" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user_answers";
DROP TABLE "question_answers";
DROP TABLE "questions";
-- +goose StatementEnd
