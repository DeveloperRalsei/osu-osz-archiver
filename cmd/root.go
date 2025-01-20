package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ooa",
	Short: "osu-osz-archiver (OOA) helps you to unzip your beatmap files in a specific directory",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	Version: "0.2.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
