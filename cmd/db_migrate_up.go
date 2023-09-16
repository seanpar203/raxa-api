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

var forceFlag = "force"
var versionFlag = "version"

// up represents the migrate up command
var dbMigrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Migrates the database all the way to the latest state.",
	Run: func(cmd *cobra.Command, args []string) {

		force, err := cmd.Flags().GetBool(forceFlag)

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		version, err := cmd.Flags().GetInt(versionFlag)

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		if force && version == 0 {
			log.Panic().Msg("You must specify a version to force")
		}

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

		if force {
			if err := m.Force(version); err != nil {
				log.Panic().Msg(err.Error())
				return
			}
		}

		if err := m.Up(); err != nil {
			switch err {
			case migrate.ErrNoChange:
				log.Info().Msg("No changes to migrate")
				return
			default:
				log.Panic().Msg(err.Error())
				return
			}
		}

		log.Info().Msg("Successfully migrated database")
	},
}

func init() {
	dbMigrateUpCmd.Flags().BoolP(forceFlag, "f", false, "Force the migration")
	dbMigrateUpCmd.Flags().IntP(versionFlag, "v", 0, "The version of the migration to force.")

	migrateCmd.AddCommand(dbMigrateUpCmd)
}
