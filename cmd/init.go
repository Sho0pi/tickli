package cmd

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

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
	Short: "Initialize tickli with your TickTick credentials",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to load config")
		}

		if clientID != "" {
			cfg.ClientID = clientID
		}
		if clientSecret != "" {
			cfg.ClientSecret = clientSecret
		}

		// Start local server for OAuth callback
		server := &http.Server{Addr: ":8080"}
		code := make(chan string)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			code <- r.URL.Query().Get("code")
			fmt.Fprintf(w, "Authorization successful! You can close this window.")
			go server.Close()
		})

		// Open browser for authorization
		authURL := api.GetAuthURL(cfg.ClientID)
		openBrowser(authURL)

		log.Info().Msg("Waiting for authorization...")
		go server.ListenAndServe()

		// Get access token
		authCode := <-code
		token, err := api.GetAccessToken(cfg.ClientID, cfg.ClientSecret, authCode)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get access token")
		}

		cfg.AccessToken = token
		if err := config.Save(cfg); err != nil {
			log.Fatal().Err(err).Msg("Failed to save config")
		}

		log.Info().Msg("Successfully initialized tickli!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&clientID, "client-id", "", "TickTick OAuth client ID")
	initCmd.Flags().StringVar(&clientSecret, "client-secret", "", "TickTick OAuth client secret")
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Error().Err(err).Msg("Failed to open browser")
	}
}
