package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "./fromstring.txt"

	content, err := ioutil.ReadFile(filename)
	checkerror(err)

	fmt.Println("Read from file: ", content)
	result := string(content)
	fmt.Println("Read from file)(string): ", resutl)
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
