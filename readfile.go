package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "./fromstring.txt"

	content, err := ioutil.ReadFile(filename)
	checkerror(err)

	fmt.Println(content)
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
