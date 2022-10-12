-- +goose Up
-- +goose StatementBegin

create table rates
(
    id integer
);

create index rates_charcode_ts_idx on rates(charcode, ts);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop index rates_code_ts_idx;
drop table rates;

-- +goose StatementEnd
