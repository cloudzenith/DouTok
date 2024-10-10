-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `comment` (
    id BIGINT PRIMARY KEY,
    video_id BIGINT NOT NULL,
    `user_id` BIGINT NOT NULL COMMENT '发表评论的用户id',
    parent_id BIGINT DEFAULT NULL COMMENT '父评论id',
    to_user_id BIGINT DEFAULT NULL COMMENT '评论所回复的用户id',
    content varchar(512) NOT NULL COMMENT '评论内容',
    first_comments json NOT NULL COMMENT '最开始的x条子评论',
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `video_id_idx` (video_id, is_deleted),
    INDEX `user_id_idx` (user_id, is_deleted),
    INDEX `create_time_idx` (create_time),
    INDEX `update_time_idx` (update_time)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `collection` (
    id BIGINT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `user_id_idx` (`user_id`, `is_deleted`),
    INDEX `create_time_idx` (`create_time`),
    INDEX `update_time_idx` (`update_time`)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `collection_video` (
    id BIGINT PRIMARY KEY,
    collection_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    video_id BIGINT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `collection_id_idx` (`collection_id`, `is_deleted`)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `favorite` (
    id BIGINT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    target_type INT NOT NULL COMMENT '点赞对象类型 1-视频 2-评论',
    target_id BIGINT NOT NULL COMMENT '点赞对象id',
    favorite_type INT NOT NULL COMMENT '点赞类型 1-点赞 2-踩',
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `user_id_idx` (`user_id`, `target_id`, `target_type`, `favorite_type`, `is_deleted`),
    INDEX `create_time_idx` (`create_time`),
    INDEX `update_time_idx` (`update_time`)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `follow` (
    id BIGINT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    target_user_id BIGINT NOT NULL COMMENT '被关注的用户id',
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `user_id_idx` (`user_id`, `target_user_id`, `is_deleted`)
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE account
ADD COLUMN `number` VARCHAR(15) NOT NULL DEFAULT '' COMMENT 'doutok号';
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE video
MODIFY COLUMN user_id BIGINT NOT NULL,
MODIFY COLUMN title VARCHAR(50) NOT NULL DEFAULT '',
MODIFY COLUMN description VARCHAR(255) NOT NULL DEFAULT '',
MODIFY COLUMN video_url VARCHAR(255) NOT NULL DEFAULT '',
MODIFY COLUMN cover_url VARCHAR(255) NOT NULL DEFAULT '',
MODIFY COLUMN like_count BIGINT NOT NULL DEFAULT 0,
MODIFY COLUMN comment_count BIGINT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE user
MODIFY COLUMN account_id BIGINT NOT NULL,
MODIFY COLUMN mobile VARCHAR(20) NOT NULL DEFAULT '',
MODIFY COLUMN email VARCHAR(50) NOT NULL DEFAULT '',
MODIFY COLUMN name VARCHAR(50) NOT NULL DEFAULT '',
MODIFY COLUMN avatar VARCHAR(255) NOT NULL DEFAULT '',
MODIFY COLUMN background_image VARCHAR(255) NOT NULL DEFAULT '',
MODIFY COLUMN signature VARCHAR(255) NOT NULL DEFAULT '',
MODIFY COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
MODIFY COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE user DROP INDEX `idx_users_email`;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE user DROP INDEX `idx_users_mobile`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comment;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS collection;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS collection_video;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS favorite;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS follow;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE account
DROP COLUMN `number`;
-- +goose StatementEnd
