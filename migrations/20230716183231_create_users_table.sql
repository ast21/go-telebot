-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id                bigserial
        primary key,
    telegram_id       varchar(20) not null
        constraint users_telegram_id_unique unique,
    first_name        varchar(255),
    last_name         varchar(255),
    created_at        timestamp(0) default CURRENT_TIMESTAMP,
    updated_at        timestamp(0) default CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
