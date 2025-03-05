package cmd

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
	"os"
	"strings"
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
				fmt.Printf("Are you sure you want to reset authentication? (y/N): ")
				reader := bufio.NewReader(os.Stdin)
				confirm, _ := reader.ReadString('\n')
				confirm = strings.TrimSpace(confirm)
				if confirm != "y" && confirm != "Y" {
					fmt.Println("Deletion aborted")
					return
				}
			}

			if err := config.DeleteToken(); err != nil {
				log.Fatal().Err(err).Msg("Failed to remove access token")
			}

			log.Info().Msg("Successfully removed access token. Running initialization...")
			token, err := initTickli()
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to initialize tickli")
			}
			log.Info().Str("token", token).Msg("Successfully initialized tickli")

		},
	}

	cmd.Flags().BoolVar(&opts.force, "force", false, "Reset authentication without confirmation")
	return cmd
}

func init() {
	RootCmd.AddCommand(newResetCommand())
}
