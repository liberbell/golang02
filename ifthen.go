package main

import "fmt"

func main() {
	// var x float64 = 0
	var result string

	if x := -42; x < 0 {
		result = "Less than Zero"
	} else if x == 0 {
		result = "Equal to Zero"
	} else {
		result = "Greater than Zero"
	}
	fmt.Println("Result: ", result)

	fmt.Println("Value of x: ", x)
}
