package beatmap

import (
	"os"

	"github.com/developerRalsei/osu-osz-archiver/utils"
	p "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var zipCmd = &cobra.Command{
	Use: "zip",
	Run: zipCommandFunc,
}

func init() {
	zipCmd.Flags().StringP(
		"file",
		"f",
		"",
		"Specify the osz file",
	)
	zipCmd.Flags().StringP(
		"out",
		"o",
		"out",
		"Specify the export location (without \"/\")",
	)
}

func zipCommandFunc(cmd *cobra.Command, args []string) {
	outLocation, err := cmd.Flags().GetString("out")
	if err != nil {
		p.Error.Printfln("%s", err.Error())
		os.Exit(1)
	}

	if outLocation == "out" {
		err := utils.CreateOutDirectory()
		if err != nil {
			p.Error.Printfln("%s", err.Error())
			os.Exit(1)
		}
	}

	utils.AskForBeatmapFolder(cmd)
	if err != nil {
		p.Error.Printfln("%s", err.Error())
		os.Exit(1)
	}
}
