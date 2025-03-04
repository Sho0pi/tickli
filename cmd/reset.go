package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

type resetOptions struct {
	force bool
}

func newResetCommand() *cobra.Command {
	opts := &resetOptions{}
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset tickli authentication",
		Long: `Reset tickli by removing the current access token and re-running the initialization process.
This is useful if you need to reauthenticate with TickTick.`,
		Run: func(cmd *cobra.Command, args []string) {
			if !opts.force {
				var confirm string
				fmt.Printf("Are you sure you want to reset authentication? (y/N): ")
				fmt.Scanln(&confirm)
				if confirm != "y" && confirm != "Y" {
					fmt.Println("Deletion aborted")
					return
				}
			}

			if err := config.DeleteToken(); err != nil {
				log.Fatal().Err(err).Msg("Failed to remove access token")
			}

			log.Info().Msg("Successfully removed access token. Running initialization...")
			initCmd.Run(cmd, args)
		},
	}

	cmd.Flags().BoolVar(&opts.force, "force", false, "Reset authentication without confirmation")
	return cmd
}

func init() {
	RootCmd.AddCommand(newResetCommand())
}
