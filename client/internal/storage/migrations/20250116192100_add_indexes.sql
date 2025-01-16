-- +goose Up
-- +goose StatementBegin

CREATE INDEX users_login_idx on users (login);
CREATE INDEX secrets_user_id_idx on secrets (user_id);
CREATE INDEX sync_events_version_idx on sync_events (version);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX users_login_idx;
DROP INDEX secrets_user_id_idx;
DROP INDEX sync_events_version_idx;
-- +goose StatementEnd
