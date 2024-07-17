-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
    id BIGINT PRIMARY KEY,
    mobile VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(64) NOT NULL,
    salt VARCHAR(64) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
-- +goose StatementEnd
