-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id                      BIGINT PRIMARY KEY,
    created_at              TIMESTAMP DEFAULT now(),
    updated_at              TIMESTAMP,
    report_abbreviation     TEXT
);

CREATE INDEX users_id_idx ON users(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS users_userid_idx;
DROP TABLE IF EXISTS users;

-- +goose StatementEnd
