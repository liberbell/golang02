package main

import "fmt"

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors.Append(colors, "Purple")
	fmt.Println(colors)
}
