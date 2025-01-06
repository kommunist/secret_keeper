-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sync_events (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    kind text,
    version text
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sync_events;
-- +goose StatementEnd
