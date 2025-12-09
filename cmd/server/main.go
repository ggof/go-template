package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ggof/template/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/knadh/koanf/parsers/toml/v2"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)


type ServerConfig struct {
	Host string
	Port uint
}

type Config struct {
	Server ServerConfig
}

func main() {
	log := zerolog.New(zerolog.NewConsoleWriter())

	k := koanf.New(".")
	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := k.Load(env.Provider("APP__", "__", parseEnv), nil); err != nil {
		log.Fatal().Err(err).Send()
	}

	var config Config

	if err := k.Unmarshal("", &config); err != nil {
		log.Fatal().Err(err).Send()
	}

	r := chi.NewRouter()

	r.Get("/", handlers.GetIndexHandler())

	hostAndPort := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	log.Info().Msgf("Starting server on port %s", hostAndPort)
	log.Fatal().Err(http.ListenAndServe(hostAndPort, r)).Send()
}

func parseEnv(key string) string {
	parts := strings.Split(key, "__")
	parts = parts[1:] // drop prefix

	upperGap := byte('a' - 'A')

	for i := range parts {
		b := strings.Builder{}

		for _, s := range strings.Split(parts[i], "_") {
			b.Grow(len(s))
			s = strings.ToLower(s)
			b.WriteByte(s[0] - upperGap)
			b.WriteString(s[1:])
		}


		parts[i] = b.String()
	}

	return strings.Join(parts, ".")
}
