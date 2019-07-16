package main

import "fmt"

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 8
	numbers[2] = 45
	numbers[3] = 101
	numbers[4] = 231
	fmt.Println(numbers)

	numbers = append(numbers, 255)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

}
