package cmd

import (
	"exrate/internal"

	"github.com/spf13/cobra"
)

var currCodesCmd = &cobra.Command{
	Use:   "codes",
	Short: "get currency codes",
	Run: func(cmd *cobra.Command, args []string) {
		internal.GetCurrCodes()
	},
}

func init() {
	exrateCmd.AddCommand(currCodesCmd)
}
