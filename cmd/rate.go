/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"exrate/internal"
	"os"

	"github.com/spf13/cobra"
)

var exrateCmd = &cobra.Command{
	Use:   "rate [",
	Short: "cli util to get exchange rates",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		internal.GetRate(args)

	},
}

func Execute() {
	err := exrateCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
