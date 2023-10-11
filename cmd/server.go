/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/seanpar203/go-api/internal/api"
	"github.com/seanpar203/go-api/internal/env"
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

		mux := http.NewServeMux()

		mux.Handle("/v1/", svc)

		if env.APP_ENV == "dev" {
			mediaHandler := http.StripPrefix(env.MEDIA_ENDPOINT, http.FileServer(http.Dir(env.MEDIA_DIR)))
			mux.Handle(env.MEDIA_ENDPOINT, mediaHandler)
		}

		log.Info().Msg("Starting our server")

		srv := &http.Server{
			Addr:         fmt.Sprintf(":%s", env.SERVER_PORT),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			Handler:      mux,
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().BoolP(migrateFlag, "m", false, "Run migrations")
}
