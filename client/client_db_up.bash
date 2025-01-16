#!/bin/bash

GOOSE_DRIVER="postgres" \
GOOSE_DBSTRING="postgresql://postgres:postgres@localhost:5435/client_for_keeper" \
GOOSE_MIGRATION_DIR="internal/storage/migrations" \
goose up