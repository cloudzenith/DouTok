-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS template (
    id BIGINT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `create_time_idx` (`create_time`),
    INDEX `update_time_idx` (`update_time`),
    INDEX `title_idx` (`title`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS template;
-- +goose StatementEnd
