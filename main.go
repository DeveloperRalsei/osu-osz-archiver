package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filePathFlag := flag.String("file", "", "Enter the path of .osz file")
	targetPathFlag := flag.String("target", "out/", "Enter the path of target file")
	flag.Parse()

	if *filePathFlag == "" {
		fmt.Println("Please enter a valid path")
		return
	}
	if !strings.HasSuffix(*filePathFlag, ".osz") {
		fmt.Println("File is not valid (osz)")
		return
	}

	fmt.Println("Opening file...")
	r, err := zip.OpenReader(*filePathFlag)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	defer r.Close()

	if *targetPathFlag == "out/" {
		mkdirIfOutPathIsNotDefined()
	}

	beatmapFolder, mathced := strings.CutSuffix(*filePathFlag, ".osz")
	if !mathced {
		panic("File could not mathced")
	}

	err = os.Mkdir(
		beatmapFolder,
		0755,
	)
	if err != nil {
		panic(err)
	}

	for _, f := range r.File {
		targetPath := filepath.Join(*targetPathFlag, f.Name)

		err = extarctFile(f, targetPath)
		if err != nil {
			fmt.Printf("Error extracting file %v: %v\n", f.Name, err)
			continue
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
