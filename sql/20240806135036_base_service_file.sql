-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS file (
    id BIGINT PRIMARY KEY,
    domain_name VARCHAR(100) NOT NULL,
    biz_name VARCHAR(100) NOT NULL,
    hash VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL DEFAULT 0,
    file_type VARCHAR(255) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `create_time_idx` (`create_time`),
    INDEX `update_time_idx` (`update_time`),
    INDEX `hash_idx` (`hash`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS file;
-- +goose StatementEnd
