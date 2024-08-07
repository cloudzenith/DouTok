-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
    id BIGINT PRIMARY KEY,
    mobile VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(64) NOT NULL,
    salt VARCHAR(64) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `account_mobile_idx` (`mobile`),
    INDEX `account_email_idx` (`email`),
    INDEX `create_time_idx` (`create_time`),
    INDEX `update_time_idx` (`update_time`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
-- +goose StatementEnd
