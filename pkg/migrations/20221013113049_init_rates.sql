-- +goose Up
-- +goose StatementBegin

CREATE TABLE rates
(
    id              TEXT PRIMARY KEY,
    created_at      TIMESTAMP DEFAULT now(),
    updated_at      TIMESTAMP,
    abbreviation    TEXT NOT NULL,
    name            TEXT,
    value           FLOAT NOT NULL
);

CREATE INDEX rates_abbreviation_idx ON rates(abbreviation);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS rates_abbreviation_idx;
DROP TABLE IF EXISTS rates;

-- +goose StatementEnd
