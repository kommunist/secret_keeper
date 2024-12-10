#!/bin/bash

GOOSE_DRIVER="postgres" \
GOOSE_DBSTRING="postgresql://postgres:postgres@localhost:5435/secret_keeper" \
GOOSE_MIGRATION_DIR="internal/client/storage/migrations" \
goose up