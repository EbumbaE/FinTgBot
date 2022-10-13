-- +goose Up
-- +goose StatementBegin

IF NOT EXISTS (SELECT * FROM rates) 
CREATE TABLE rates
(
    created_at      TIMESTAMPTZ NOT NULL DEFAULT (now()),
    abbreviation    VARCHAR NOT NULL,
    name            VARCHAR NOT NULL,
    value           FLOAT NOT NULL,
);

IF NOT EXISTS (SELECT * FROM users) 
CREATE TABLE users
(
    created_at      TIMESTAMPTZ NOT NULL DEFAULT (now()),
    user_id         BIGSERIAL PRIMARY KEY,
    abbreviation    VARCHAR NOT NULL,
);

IF NOT EXISTS (SELECT * FROM diary) 
CREATE TABLE diary
(
    created_at      TIMESTAMPTZ NOT NULL DEFAULT (now()),
    user_id         BIGSERIAL PRIMARY KEY,
    date            VARCHAR NOT NULL,
    note_category   VARCHAR NOT NULL,
    note_currency   VARCHAR NOT NULL,
    note_sum        FLOAT NOT NULL,
);

CREATE INDEX users_userid_ts_idx ON users(user_id);
CREATE INDEX diary_userid_date_idx ON diary(user_id, date);
CREATE INDEX users_userid_ts_idx ON users(user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS rates_abbreviation_ts_idx;
DROP TABLE IF EXISTS rates;

DROP INDEX IF EXISTS users_userid_ts_idx;
DROP TABLE IF EXISTS  users;

DROP INDEX IF EXISTS diary_userid_date_idx;
DROP TABLE IF EXISTS diary;

-- +goose StatementEnd
