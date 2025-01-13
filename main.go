package main

import (
	"archive/zip"
	"fmt"
	"os"
	"strings"
)

// default flags

var out string = "out/"

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Please enter a parameter first\n")
		return
	}

	filePath := args[1]

	if !strings.HasSuffix(filePath, ".osz") {
		fmt.Printf("The file format must be \"osz\" \n")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := zip.NewWriter(file)
	defer writer.Close()

	directories, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	exists := false
	for _, dir := range directories {
		if dir.Name() == "out" && dir.IsDir() {
			exists = true
			break
		}
	}
	if !exists {
		os.Mkdir("out", 0755)
	}

	beatmapFolder := strings.SplitAfter(
		fmt.Sprintf("out/%s", file.Name()),
		".osz",
	)

	fmt.Printf("beatmapFolder: %v\n", beatmapFolder)
}
