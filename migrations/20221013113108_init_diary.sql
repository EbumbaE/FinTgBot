-- +goose Up
-- +goose StatementBegin

CREATE TABLE diary
(
    created_at      TIMESTAMP NOT NULL,
    user_id         BIGSERIAL,
    date            TEXT NOT NULL,
    note_category   TEXT NOT NULL,
    note_currency   TEXT NOT NULL,
    note_sum        FLOAT NOT NULL
);

CREATE INDEX diary_userid_date_idx ON diary(user_id, date);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS diary_userid_date_idx;
DROP TABLE IF EXISTS diary;

-- +goose StatementEnd
