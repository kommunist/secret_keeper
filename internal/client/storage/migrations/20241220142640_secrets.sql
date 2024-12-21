-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE secrets (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name text,
    pass text,
    meta text,

    version text,

    user_id uuid REFERENCES users (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE secrets;
-- +goose StatementEnd
