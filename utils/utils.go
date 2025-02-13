package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	p "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func CreateOutDirectory() error {
	stat, err := os.Stat("out")
	if err == nil && stat.IsDir() {
		return nil
	}

	if os.IsNotExist(err) {
		return os.Mkdir("out", 0755)
	}

	return err
}

// Asks user for beatmap file via cmd
func AskForBeatmapFileViaCmd(cmd *cobra.Command) (*os.File, error) {
	beatmap, error := cmd.Flags().GetString("file")
	if error != nil {
		p.Error.Printfln("Error: %s", error.Error())
		return nil, error
	}

	textInput := p.DefaultInteractiveTextInput.WithDefaultText("Write path of beatmap file")

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

		valid, errMsg := CheckOSZFile(beatmap_file)
		if !valid {
			p.Error.Println(errMsg)
			beatmap = ""
			beatmap_file.Close()
			continue
		}

		return beatmap_file, nil

	}
}

func AskForBeatmapFolder(cmd *cobra.Command) {
	textInput := p.DefaultInteractiveTextInput.WithDefaultText("Write path of beatmap folder")
	_, err := cmd.Flags().GetString("file")
	if err != nil {
		p.Error.Println(err.Error())
		os.Exit(1)
	}

	for {
		textInput.Show("")
	}
}

// Checks if the file complies with the osz type and returns a bool and a error message
func CheckOSZFile(file *os.File) (bool, string) {
	fileInfo, err := file.Stat()
	if err != nil {
		return false, fmt.Sprintf(
			"Error while gettings file stats: %s",
			err.Error(),
		)
	}

	if fileInfo.IsDir() {
		return false, "You choose a dir"
	}

	if !strings.HasSuffix(
		file.Name(),
		".osz",
	) {
		return false, "The file is not a osz file type"
	}

	return true, ""
}

/*
Creates a folder that holding beatmap files and returns the existed/created folder path. Takes to arguments.
1. file: beatmap file
2. where: spesific folder that containes beatmap file
*/
func CreateBeatmapFolder(file *os.File, where string) (string, error) {
	beatmapFolderName := filepath.Base(
		strings.TrimSuffix(file.Name(), ".osz"),
	)

	entries, err := os.ReadDir(where)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.Name() == beatmapFolderName && entry.IsDir() {
			return filepath.Join(where, beatmapFolderName), nil
		}
	}

	folderPath := filepath.Join(
		where,
		beatmapFolderName,
	)

	if err := os.Mkdir(folderPath, 0755); err != nil {
		return "", err
	}
	p.Info.Println("Out directory created")
	return filepath.Join(folderPath), nil
}
