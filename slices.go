package main

import "fmt"

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[1:])
	fmt.Println(colors)
}
