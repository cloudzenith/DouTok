-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
    id BIGINT PRIMARY KEY,
    mobile VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(64) NOT NULL,
    salt VARCHAR(64) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    INDEX `account_mobile_idx` (mobile),
    INDEX `account_email_idx` (email)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `account_id` bigint(20) DEFAULT NULL,
    `mobile` varchar(20) DEFAULT NULL,
    `email` varchar(50) DEFAULT NULL,
    `name` varchar(50) DEFAULT NULL,
    `avatar` varchar(255) DEFAULT NULL,
    `background_image` varchar(255) DEFAULT NULL,
    `signature` varchar(255) DEFAULT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_users_account_id` (`account_id`) USING BTREE,
    UNIQUE KEY `idx_users_email` (`email`) USING BTREE,
    UNIQUE KEY `idx_users_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7240204103809514232 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
-- +goose StatementEnd

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

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS file (
    id BIGINT PRIMARY KEY,
    domain_name VARCHAR(100) NOT NULL,
    biz_name VARCHAR(100) NOT NULL,
    hash VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL DEFAULT 0,
    file_type VARCHAR(255) NOT NULL,
    uploaded BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `create_time_idx` (`create_time`),
    INDEX `update_time_idx` (`update_time`),
    INDEX `hash_idx` (`hash`)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE `video` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `user_id` bigint(20) DEFAULT NULL,
     `title` varchar(20) DEFAULT NULL,
     `description` varchar(50) DEFAULT NULL,
     `video_url` varchar(255) DEFAULT NULL,
     `cover_url` varchar(255) DEFAULT NULL,
     `like_count` bigint(20) DEFAULT 0,
     `comment_count` bigint(20) DEFAULT 0,
     `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `user`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS video;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS template;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS file;
-- +goose StatementEnd
