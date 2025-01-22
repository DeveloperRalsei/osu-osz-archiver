/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/developerRalsei/osu-osz-archiver/osz"
	p "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var textInput = p.DefaultInteractiveTextInput.WithDefaultText("Write path of beatmap file")

var unzipCommand = &cobra.Command{
	Use: "unzip",
	Run: unzipCommandFunc,
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

func unzipCommandFunc(cmd *cobra.Command, args []string) {
	outLocation, err := cmd.Flags().GetString("out")
	if err != nil {
		p.Error.Printfln("%s 21", err.Error())
		os.Exit(1)
	}

	if outLocation == "out" {
		err := createOutDirectory()
		if err != nil {
			p.Error.Printfln("%s 28", err.Error())
			os.Exit(1)
		}
	}

	beatmap_file := askForBeatmapFileViaCmd(cmd)
	defer beatmap_file.Close()

	beatmap_folder, err := osz.CreateBeatmapFolder(beatmap_file, outLocation)
	if err != nil {
		p.Error.Printfln("%s 38", err.Error())
		os.Exit(1)
	}

	fmt.Printf("beatmap_folder: %v\n", beatmap_folder)
}

func init() {
	beatmapCmd.AddCommand(unzipCommand)
	unzipCommand.Flags().StringP("file", "f", "", "Specify the osz file")
	unzipCommand.Flags().StringP("out", "o", "out", "Specify the export location (without \"/\"")
}
