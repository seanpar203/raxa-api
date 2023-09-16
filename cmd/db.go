/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db <subcommand>",
	Short: "Entrypoint for doing operations on the database",
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
