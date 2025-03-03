package task

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

var TickliClient *api.Client

var (
	projectID string
)

var Cmd = &cobra.Command{
	Use:   "task",
	Short: "Work on TickTick tasks.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		token, err := config.LoadToken()
		if err != nil || token == "" {
			log.Fatal().Msg("Please run 'tickli init' first")
		}

		// Init the TickliClient
		TickliClient = api.NewClient(token)

		if projectID == "" {
			cfg, err := config.Load()
			if err != nil {
				return errors.Wrap(err, "failed to load config")
			}
			projectID = cfg.DefaultProjectID
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
	},
}

func init() {
	Cmd.AddCommand(newCreateCommand())
	Cmd.AddCommand(newUpdateCommand())
	Cmd.AddCommand(newListCommand())
	Cmd.AddCommand(newCompleteCmd())
	Cmd.AddCommand(newUncompleteCommand())
	Cmd.AddCommand(newDeleteCommand())
	Cmd.AddCommand(newShowCommand())
}
