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

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs our server",
	Run: func(cmd *cobra.Command, args []string) {

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
}
