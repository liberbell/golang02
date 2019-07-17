package main

import "fmt"

func main() {
	var x float64 = 42
	var result string

	if x < 0 {
		result = "Less than Zero"
	} else if x == 0 {
		result = "Equal to Zero"
	} else {
		result = "Greater than Zero"
	}
	fmt.Println("Result: ", result)
}
