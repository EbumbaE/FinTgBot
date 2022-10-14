-- +goose Up
-- +goose StatementBegin

CREATE TABLE budgets
(
    created_at      TIMESTAMP NOT NULL,
    user_id         BIGINT,
    date            TEXT PRIMARY KEY,
    value           FLOAT NOT NULL,
    abbreviation    TEXT NOT NULL
);

CREATE INDEX budgets_userid_date_idx ON budgets(user_id, date);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS budgets_userid_date_idx;
DROP TABLE IF EXISTS budgets;

-- +goose StatementEnd
