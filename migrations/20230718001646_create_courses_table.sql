-- +goose Up
-- +goose StatementBegin
CREATE TABLE "courses" (
    "id" bigserial PRIMARY KEY,
    "title" varchar,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_has_courses" (
    "user_id" bigserial,
    "course_id" bigserial
);

ALTER TABLE "user_has_courses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_has_courses" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user_has_courses";
DROP TABLE "courses";
-- +goose StatementEnd
