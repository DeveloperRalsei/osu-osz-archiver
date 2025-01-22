package cmd

import (
	"github.com/spf13/cobra"
)

var beatmapCmd = &cobra.Command{
	Use:   "beatmap",
	Short: "Define an osz file with file flag",
	Long:  "Zip and Unzip beatmap files",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(beatmapCmd)
}
