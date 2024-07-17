-- +goose Up
-- +goose StatementBegin
ALTER TABLE doutok.account ADD INDEX `account_mobile_idx` (`mobile`), ADD INDEX `account_email_idx` (`email`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE doutok.account DROP INDEX account_mobile_idx, DROP INDEX account_email_idx;
-- +goose StatementEnd
