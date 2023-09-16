/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate <subcommand>",
	Short: "Entrypoint for performing specific types of migrations on the database.",
}

func init() {
	dbCmd.AddCommand(migrateCmd)
}
