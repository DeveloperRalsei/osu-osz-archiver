package cmd

import (
	"os"

	"github.com/developerRalsei/osu-osz-archiver/osz"
	p "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var beatmapCmd = &cobra.Command{
	Use:   "beatmap",
	Short: "Define an osz file with file flag",
	Run: func(cmd *cobra.Command, args []string) {
		beatmap, _ := cmd.Flags().GetString("file")
		textInput := p.DefaultInteractiveTextInput.WithDefaultText("Write path of beatmap file")

	openFileViaTextInput:
		var beatmap_file *os.File
		if beatmap == "" {
			value, _ := textInput.Show("")
			beatmap = value
		}

		file, err := os.Open(beatmap)
		if err != nil {
			p.Error.Printfln("Error: %s", err.Error())
			beatmap = ""
			goto openFileViaTextInput

		}
		defer file.Close()
		beatmap_file = file

		valid, errMsg := osz.CheckOSZFile(beatmap_file)
		if !valid {
			p.Error.Println(errMsg)
			beatmap = ""
			goto openFileViaTextInput
		}

		out, err := cmd.Flags().GetString("out")
		if err != nil {
			p.Error.Printfln("Something went wrong while getting out flag value : %s", err.Error())
			os.Exit(1)
		}

		if out == "out/" {
			entries, err := os.ReadDir(".")
			if err != nil {
				p.Error.Printfln("Error: %s", err.Error())
			}

			isOutDirExist := false
			for _, entry := range entries {
				if entry.IsDir() && entry.Name() == "out/" {
					isOutDirExist = true
					break
				}
			}

			if !isOutDirExist {
				err = os.Mkdir("out", 0755)
				if err != nil {
					p.Error.Printfln("Something went wrong when trying to create out folder: %s", err.Error())
				}
				p.Info.Println("Created an out directory")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(beatmapCmd)
	beatmapCmd.Flags().StringP("file", "f", "", "Specify the osz file")
	beatmapCmd.Flags().StringP("out", "o", "out/", "Specify the export location")
}
