/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"exrate/internal"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var exrateCmd = &cobra.Command{
	Use:   "exrate",
	Short: "cli util to get exchange rates",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		internal.GetRate(strings.ToUpper(args[0]))

	},
}

func Execute() {
	err := exrateCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.exrate.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	exrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
