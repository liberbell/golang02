package main

import "fmt"

func main() {
	dosomething()

	sum := AddValues(22, 11)
	fmt.Println("Value: ", sum)

	sum = AddAllValues(12, 54, 59)
	fmt.Println("Value: ", sum)
}

func dosomething() {
	fmt.Println("Doing something!")
}

func AddValues(value1 int, value2 int) int {
	return value1 + value2
}

func AddAllValues(values ...int) int {
	sum := 0
	fmt.Printf("%T\n", values)
	return sum
}
