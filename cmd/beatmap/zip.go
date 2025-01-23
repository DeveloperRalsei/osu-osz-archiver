package beatmap

import (
	"fmt"

	"github.com/spf13/cobra"
)

var zipCmd = &cobra.Command{
	Use: "zip",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zip called")
	},
}

func init() {
	BeatmapCmd.AddCommand(zipCmd)
}
