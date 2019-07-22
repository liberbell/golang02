package main

import (
	"io"
	"os"
)

func main() {
	content := "Hello from Go."
	file, err := os.Create("./fromstring.txt")
	checkerror(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkerror(err)
}

func checkerror() {
	if err != nil {
		panic(err)
	}
}
