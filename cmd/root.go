package cmd

import (
	"os"

	"github.com/developerRalsei/osu-osz-archiver/cmd/beatmap"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ooa",
	Short: "osu-osz-archiver (OOA) helps you to unzip your beatmap files in a specific directory",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	Version: "0.2.0",
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(beatmap.BeatmapCmd)
}
