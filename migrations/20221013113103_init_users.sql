-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    created_at              TIMESTAMP NOT NULL,
    user_id                 BIGINT PRIMARY KEY,
    report_abbreviation     TEXT NOT NULL
);

CREATE INDEX users_userid_idx ON users(user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS users_userid_idx;
DROP TABLE IF EXISTS users;

-- +goose StatementEnd
