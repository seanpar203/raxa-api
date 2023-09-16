/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/seanpar203/go-api/internal/db"
)

// dbMigrateDownCmd represents the dbMigrateDown command
var dbMigrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rolls back all of the migrations.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.Postgres()

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		driver, err := postgres.WithInstance(db, &postgres.Config{})

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		err = m.Down()

		if err != nil {
			switch err {
			case migrate.ErrNoChange:
				log.Info().Msg("No changes to roll back")
				return
			default:
				log.Panic().Msg(err.Error())
			}
		}

		log.Info().Msg("Successfully rolled back database")
	},
}

func init() {
	migrateCmd.AddCommand(dbMigrateDownCmd)
}
