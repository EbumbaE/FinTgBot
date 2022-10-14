-- +goose Up
-- +goose StatementBegin

CREATE TABLE rates
(
    created_at      TIMESTAMP NOT NULL,
    abbreviation    TEXT NOT NULL,
    name            TEXT NOT NULL,
    value           FLOAT NOT NULL
);

CREATE INDEX rates_abbreviation_idx ON rates(abbreviation);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS rates_abbreviation_idx;
DROP TABLE IF EXISTS rates;

-- +goose StatementEnd
