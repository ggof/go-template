#! /usr/bin/env bash

GOOSE_DRIVER=sqlite3
GOOSE_DBSTRING="./project.db"
GOOSE_MIGRATION_DIR="db/migrations"

go tool goose -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER $GOOSE_DBSTRING $@
