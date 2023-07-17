-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "telegram_id" varchar(20) UNIQUE NOT NULL,
    "first_name" varchar,
    "last_name" varchar,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd
