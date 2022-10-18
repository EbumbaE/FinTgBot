-- +goose Up
-- +goose StatementBegin

CREATE TABLE budgets
(
    user_id         BIGINT PRIMARY KEY,
    created_at      TIMESTAMP DEFAULT now(),
    updated_at      TIMESTAMP,
    date            TEXT,
    value           FLOAT,
    abbreviation    TEXT
);

CREATE INDEX budgets_userid_date_idx ON budgets(user_id, date);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS budgets_userid_date_idx;
DROP TABLE IF EXISTS budgets;

-- +goose StatementEnd
