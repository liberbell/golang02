package main

import "fmt"

func main() {
	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break
		}
	}

endofprogram:
	fmt.Println("End of Program")
}
