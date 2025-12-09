#! /usr/bin/env bash

alias wgo="go tool wgo"
alias templ="go tool templ"
alias sqlc="go tool sqlc"

# start by running the migrations
./goose.sh up

wgo -xdir=db/queries -xfile=_templ.go templ generate :: sqlc generate :: go run ./cmd/server/main.go

./goose.sh down

unalias wgo
unalias templ
unalias sqlc
