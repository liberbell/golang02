package main

import "os"

func main() {
	content := "Hello from Go."
	file, err := os.Create("./fromstring.txt")
	checkerror(err)
	defer file.Close()
}

func checkerror() {
	if err != nil {
		panic(err)
	}
}
