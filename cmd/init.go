package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
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
	Short: "Initialize tickli and obtain an access token",
	Long: `Initialize tickli by performing OAuth authentication with TickTick.
This will open your browser for authentication and store the access token securely.`,
	Run: func(cmd *cobra.Command, args []string) {

		if token, err := config.LoadToken(); err != nil {
			log.Fatal().Err(err).Msg("failed to check existing token")
		} else if token != "" {
			log.Fatal().Msg("tickli is already initialized. Used 'tickli reset' to reinitialize")
		}

		if err := godotenv.Load(); err == nil {
			log.Info().Msg("Loading TickTick credentials from .env")

			clientID = os.Getenv("TICKTICK_CLIENT_ID")
			clientSecret = os.Getenv("TICKTICK_CLIENT_SECRET")
		}

		// Verify credentials are available
		if clientID == "" || clientSecret == "" {
			log.Fatal().Msg("Missing TickTick credentials. Please provide them via environment variables or build flags")
		}

		// Start OAuth flow
		server := &http.Server{Addr: ":8080"}
		code := make(chan string)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			code <- r.URL.Query().Get("code")
			fmt.Fprintf(w, "Authorization successful! You can close this window.")
			go server.Close()
		})

		authURL := api.GetAuthURL(clientID)
		if err := openBrowser(authURL); err != nil {
			log.Fatal().Err(err).Msg("Failed to open browser")
		}

		log.Info().Msg("Waiting for authorization...")
		go server.ListenAndServe()

		// Get access token
		authCode := <-code
		token, err := api.GetAccessToken(clientID, clientSecret, authCode)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get access token")
		}

		if err := config.SaveToken(token); err != nil {
			log.Fatal().Err(err).Msg("Failed to save access token")
		}

		log.Info().Msg("Successfully initialized tickli!")
	},
}

func openBrowser(url string) (err error) {
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func init() {
	RootCmd.AddCommand(initCmd)
}
