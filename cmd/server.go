/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/seanpar203/go-api/internal/api"
)

var migrateFlag = "migrate"

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs our server",
	Run: func(cmd *cobra.Command, args []string) {

		migrate, err := cmd.Flags().GetBool(migrateFlag)

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		if migrate {
			dbMigrateUpCmd.Run(cmd, args)
		}

		svc, err := api.New()

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		log.Info().Msg("Starting our server")

		if err := http.ListenAndServe(":8080", svc); err != nil {
			log.Fatal().Msg(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().BoolP(migrateFlag, "m", false, "Run migrations")
}
