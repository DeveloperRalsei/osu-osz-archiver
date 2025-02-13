package beatmap

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/developerRalsei/osu-osz-archiver/utils"
	p "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var unzipCommand = &cobra.Command{
	Use: "unzip",
	Run: unzipCommandFunc,
}

func init() {
	unzipCommand.Flags().StringP("file", "f", "", "Specify the osz file")
	unzipCommand.Flags().StringP("out", "o", "out", "Specify the export location (without \"/\"")
}

func unzipCommandFunc(cmd *cobra.Command, args []string) {
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

	beatmap_file, err := utils.AskForBeatmapFileViaCmd(cmd)
	if err != nil {
		p.Error.Printfln("%s", err.Error())
		os.Exit(1)
	}
	defer beatmap_file.Close()

	beatmap_file_path := filepath.Join(beatmap_file.Name())

	beatmap_folder, err := utils.CreateBeatmapFolder(beatmap_file, outLocation)
	if err != nil {
		p.Error.Printfln("%s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("beatmap_folder: %v\n", beatmap_folder)

	r, err := zip.OpenReader(beatmap_file_path)
	if err != nil {
		p.Error.Printfln("%s", err.Error())
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

		destFile, err := os.Create(destPath)
		if err != nil {
			p.Error.Printfln("Failed to create destination file: %s", err.Error())
			srcFile.Close()
			continue
		}

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			p.Error.Printfln("Failed to copy file contents: %s", err.Error())
		}
		p.Success.Printfln("🎊️ Copied file successfuly: %s", destFile.Name())
	}
}
