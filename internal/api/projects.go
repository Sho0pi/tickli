package api

import (
	"fmt"
	"github.com/sho0pi/tickli/internal/config"
)

func GetClient() (*Client, error) {
	token, err := config.LoadToken()
	if err != nil {
		return nil, fmt.Errorf("falied to load token: %w", err)
	}

	return NewClient(token), nil
}
