package cmd

import (
	"os"

	"github.com/developerRalsei/osu-osz-archiver/osz"
	p "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var textInput = p.DefaultInteractiveTextInput.WithDefaultText("Write path of beatmap file")

var beatmapCmd = &cobra.Command{
	Use:   "beatmap",
	Short: "Define an osz file with file flag",
	Run: func(cmd *cobra.Command, args []string) {
		outLocation, err := cmd.Flags().GetString("out")
		if err != nil {
			p.Error.Printfln("%s 19", err.Error())
			os.Exit(1)
		}

		if outLocation == "out" {
			err := createOutDirectory()
			if err != nil {
				p.Error.Printfln("%s 26", err.Error())
				os.Exit(1)
			}
		}

		beatmap_file := askForBeatmapFileViaCmd(cmd)
		defer beatmap_file.Close()

		err = osz.CreateBeatmapFolder(beatmap_file, outLocation)
		if err != nil {
			p.Error.Printfln("%s 36", err.Error())
			os.Exit(1)
		}
	},
}

// Asks user for beatmap file via cmd
func askForBeatmapFileViaCmd(cmd *cobra.Command) *os.File {
	beatmap, _ := cmd.Flags().GetString("file")

	for {
		if beatmap == "" {
			value, _ := textInput.Show("")
			beatmap = value
		}

		beatmap_file, err := os.Open(beatmap)
		if err != nil {
			p.Error.Printfln("%s 54", err.Error())
			beatmap = ""
			continue
		}

		valid, errMsg := osz.CheckOSZFile(beatmap_file)
		if !valid {
			p.Error.Println(errMsg)
			beatmap = ""
			beatmap_file.Close()
			continue
		}

		return beatmap_file

	}
}

func createOutDirectory() error {
	stat, err := os.Stat("out")
	if err == nil && stat.IsDir() {
		return nil
	}

	if os.IsNotExist(err) {
		return os.Mkdir("out", 0755)
	}

	return err
}

func init() {
	rootCmd.AddCommand(beatmapCmd)
	beatmapCmd.Flags().StringP("file", "f", "", "Specify the osz file")
	beatmapCmd.Flags().StringP("out", "o", "out", "Specify the export location (without \"/\"")
}
