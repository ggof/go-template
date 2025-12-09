# Go Template

This repository is a template for Go HTTP servers. It currently provides my own very opinionated way of doing things. With this project, you will use:

- [Chi](https://github.com/go-chi/chi) for HTTP routing
- [Sqlc](https://sqlc.dev/) for database interactions (with SQLite currently configured)
- [Goose](https://github.com/pressly/goose) for database migrations (with SQLite currently configured)
- [Templ](https://templ.guide/) for HTML templating
- [Koanf](https://github.com/knadh/koanf) for configuration 
- [Wgo](https://github.com/bokwoon95/wgo) for hot-ish reloading

## Tool Configuration
For the tools (`wgo`, `templ`, `goose` and `sqlc`), most of the tools will be available via `go tool <the tool>`. A minimal, working configuration for each of those tools is provided.

N.B: `goose` requires configuration to run properly, so I've added a simple `goose.sh` script to wrap the executable and provide the configuration. be sure to update them for your setup!

## Application Configuration
Application config is done through `koanf` and the format is up to you. What is configured in this repo is TOML files + environment, but you can definitely add some dev files etc.

## Entrypoint
The entrypoints are located in packages under `cmd`. For instance, the HTTP server lives under `cmd/server/main.go`
