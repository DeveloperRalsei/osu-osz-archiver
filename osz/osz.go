package osz

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	p "github.com/pterm/pterm"
)

// func ReadAndExportOSZ(file *os.File) error {
// }

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
