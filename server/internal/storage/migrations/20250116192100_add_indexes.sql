-- +goose Up
-- +goose StatementBegin

CREATE INDEX users_login_idx on users (login);
CREATE INDEX secrets_user_id_version_idx on secrets (user_id, version);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX users_login_idx;
DROP INDEX secrets_user_id_idx;
-- +goose StatementEnd
