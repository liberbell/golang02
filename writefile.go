package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go."
	file, err := os.Create("./fromstring.txt")
	checkerror(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkerror(err)

	fmt.Printf("all done with file of %v characters.\n", ln)

	bytes := []byte(content)
	ioutil.WriteFile("./fromBytes.txt", bytes)
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
