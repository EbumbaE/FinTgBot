-- +goose Up
-- +goose StatementBegin

CREATE TABLE diary
(
    user_id         BIGSERIAL,
    created_at      TIMESTAMP DEFAULT now(),
    updated_at      TIMESTAMP,
    date            TEXT,
    note_category   TEXT,
    note_currency   TEXT,
    note_sum        FLOAT
);

CREATE INDEX diary_userid_date_idx ON diary(user_id, date);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS diary_userid_date_idx;
DROP TABLE IF EXISTS diary;

-- +goose StatementEnd
