package main

import "fmt"

func main() {
	dosomething()

	sum := AddValues(22, 11)
	fmt.Println("Value: ", sum)
}

func dosomething() {
	fmt.Println("Doing something!")
}

func AddValues(value1 int, value2 int) int {
	return value1 + value2
}
