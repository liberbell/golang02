package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, _ := filepath.Abs(".")
	fmt.Println("processing path", root)

	err := filepath.Walk(root, processPath)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func processPath(path string, info os.FileInfo, err error) {
	if err != nil {
		return err
	}

	if path != "." {
		if info.isDir() {
			fmt.Println("Directory: ", path)
		} else {
			fmt.Println("File: ", path)
		}
	}
	return nil
}
