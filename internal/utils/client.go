package utils

import (
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
)

func LoadClient() api.Client {
	token, err := config.LoadToken()
	if err != nil {
		log.Fatal().Err(err).Msg("Please run 'tickli init' first")
	}
	return *api.NewClient(token)
}
