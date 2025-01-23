package beatmap

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

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
			p.Error.Printfln("%s 36", err.Error())
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
		p.Error.Printfln("%s 70", err.Error())
		os.Exit(1)
	}

	if outLocation == "out" {
		err := createOutDirectory()
		if err != nil {
			p.Error.Printfln("%s 77", err.Error())
			os.Exit(1)
		}
	}

	beatmap_file := askForBeatmapFileViaCmd(cmd)
	defer beatmap_file.Close()

	beatmap_file_path := filepath.Join(beatmap_file.Name())

	beatmap_folder, err := osz.CreateBeatmapFolder(beatmap_file, outLocation)
	if err != nil {
		p.Error.Printfln("%s 89", err.Error())
		os.Exit(1)
	}
	fmt.Printf("beatmap_folder: %v\n", beatmap_folder)

	r, err := zip.OpenReader(beatmap_file_path)
	if err != nil {
		p.Error.Printfln("%s 85", err.Error())
		os.Exit(1)
	}
	defer r.Close()

	for _, f := range r.File {
		destPath := filepath.Join(beatmap_folder, f.Name)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(destPath, os.ModePerm)
			if err != nil {
				p.Error.Printfln("Failed to create directory: %s", err.Error())
			}
			continue
		}

		srcFile, err := f.Open()
		if err != nil {
			p.Error.Printfln("Failed to open file in archive: %s", err.Error())
			continue
		}
		defer srcFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			p.Error.Printfln("Failed to create destination file: %s", err.Error())
			continue
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			p.Error.Printfln("Failed to copy file contents: %s", err.Error())
		}
		p.Success.Printfln("🎊️ Copied file successfuly: %s", destFile.Name())
	}
}

func init() {
	BeatmapCmd.AddCommand(unzipCommand)
	unzipCommand.Flags().StringP("file", "f", "", "Specify the osz file")
	unzipCommand.Flags().StringP("out", "o", "out", "Specify the export location (without \"/\"")
}
