package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

var (
	clientID     string
	clientSecret string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize tickli and obtain an access token",
	Long: `Initialize tickli by performing OAuth authentication with TickTick.
This will open your browser for authentication and store the access token securely.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if token, err := config.LoadToken(); err != nil {
			log.Fatal().Err(err).Msg("failed to check existing token")
		} else if token != "" {
			log.Fatal().Msg("tickli is already initialized. Used 'tickli reset' to reinitialize")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		token, err := initTickli()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to initialize tickli")
		}
		log.Info().Str("token", token).Msg("Successfully initialized tickli!")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}

func initTickli() (string, error) {
	if err := godotenv.Load(); err == nil {
		log.Info().Msg("Loading TickTick credentials from .env")

		clientID = os.Getenv("TICKTICK_CLIENT_ID")
		clientSecret = os.Getenv("TICKTICK_CLIENT_SECRET")
	}

	// Verify credentials are available
	if clientID == "" || clientSecret == "" {
		return "", fmt.Errorf("missing TickTick credentials. Please provide them via environment variables or build flags")
	}

	// Start OAuth flow
	server := &http.Server{Addr: ":8080"}
	code := make(chan string, 1)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code <- r.URL.Query().Get("code")
		fmt.Fprintf(w, "Authorization successful! You can close this window.")
		go func() {
			if err := server.Close(); err != nil {
				log.Error().Err(err).Msg("Failed to close server")
			}
		}()
	})

	authURL := api.GetAuthURL(clientID)
	if err := browser.OpenURL(authURL); err != nil {
		return "", errors.Wrap(err, "failed to open browser")
	}

	log.Info().Msg("Waiting for authorization...")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Server error")
		}
	}()

	// Get access token
	authCode := <-code
	token, err := api.GetAccessToken(clientID, clientSecret, authCode)
	if err != nil {
		return "", errors.Wrap(err, "failed to get access token")
	}

	if err := config.SaveToken(token); err != nil {
		return "", errors.Wrap(err, "failed to save token")
	}

	return token, nil
}
