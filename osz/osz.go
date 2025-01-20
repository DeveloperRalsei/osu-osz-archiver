package osz

import (
	"fmt"
	"os"
	"strings"
)

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
