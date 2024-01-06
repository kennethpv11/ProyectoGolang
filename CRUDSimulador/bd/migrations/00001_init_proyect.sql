-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SCHEMA col_electricity_info;
CREATE TABLE col_electricity_info.metric(
    id VARCHAR PRIMARY KEY,
    measure VARCHAR,
    value_max INTEGER,
    value_min INTEGER
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE col_electricity_info.metric;
DROP SCHEMA col_electricity_info;
-- +goose StatementEnd
