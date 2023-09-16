/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/seanpar203/go-api/internal/db"
)

// dbMigrateUpCmd represents the dbMigrateUp command
var dbMigrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Migrates the database all the way to the latest state.",
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

		err = m.Up()

		if err != nil {
			switch err {
			case migrate.ErrNoChange:
				log.Info().Msg("No changes to migrate")
				return
			default:
				log.Panic().Msg(err.Error())
			}
		}

		log.Info().Msg("Successfully migrated database")
	},
}

func init() {
	migrateCmd.AddCommand(dbMigrateUpCmd)

}
