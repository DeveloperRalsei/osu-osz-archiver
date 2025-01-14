package main

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filePathFlag := flag.String("file", "", "Enter the path of .osz file")

	homeDir, exist := os.LookupEnv("HOME")
	if !exist {
		fmt.Printf("HOME enviroment not set")
		os.Exit(1)
	}
	targetFlag := flag.String(
		"target",
		filepath.Join(homeDir, ".osu", "Songs"),
		"Excarting file location",
	)
	flag.Parse()

	if *filePathFlag == "" {
		fmt.Println("Please enter a valid path")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !strings.HasSuffix(*filePathFlag, ".osz") {
		fmt.Println("File is not valid (osz)")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Opening file...")
	r, err := zip.OpenReader(*filePathFlag)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	defer r.Close()

	beatmapFolderName := strings.TrimSuffix(*filePathFlag, ".osz")
	exportLocation := *targetFlag + beatmapFolderName

	// _, err := mkdirBeatmapFolder(&exportLocation, *targetFlag, homeDir)
	bFolderLocation, err := mkdirBeatmapFolder(&exportLocation, *targetFlag, homeDir)
	if err != nil {
		panic(err)
	}

	for _, f := range r.File {
		targetPath := filepath.Join(bFolderLocation, f.Name)

		err = extarctFile(f, targetPath)
		if err != nil {
			fmt.Printf("Error extracting file %v: %v\n", f.Name, err)

			break
		}
	}
}

func mkdirIfOutPathIsNotDefined() {
	dirs, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() && dir.Name() == "out/" {
			os.Mkdir("out", 0755)
			break
		}
	}
}

func extarctFile(f *zip.File, targetPath string) error {
	srcFile, err := f.Open()
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	if !f.FileInfo().IsDir() {
		err = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
		if err != nil {
			return err
		}
	}

	dstFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	fmt.Printf("Extracting %v to %v\n", f.Name, targetPath)
	_, err = io.Copy(dstFile, srcFile)
	return err
}

func mkdirBeatmapFolder(bFolder *string, targetPath string, homedir string) (beatmapFolderLocation string, err error) {
	files, err := os.ReadDir(targetPath)

	var bFolderLocName string

	for _, file := range files {
		if file.Name() != *bFolder {
			fileLocation := filepath.Join(homedir, file.Name())

			fmt.Printf("fileLocation: %v\n", fileLocation)
			bFolderLocName = file.Name()
			return fileLocation, nil
		} else {
			err := os.Mkdir(
				*bFolder,
				0755,
			)
			if err != nil {
				return "", errors.New("Couldn't create beatmapFolder")
			}
		}
	}
	return bFolderLocName, nil
}
