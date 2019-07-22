package main

import "fmt"

func main() {
	defer fmt.Println("Close the file.")
	fmt.Println("Open the file.")

	defer fmt.Println("Statement 1.")
	defer fmt.Println("Statement 2.")

	myFunc()

	defer fmt.Println("Statement 3.")
	defer fmt.Println("Statement 4.")
	fmt.Println("Undefer Statement.")

	x := 1000
	defer fmt.Println("value of x:", x)
	x++
	fmt.Println("value of x:", x)
}

func myFunc() {
	defer fmt.Println("Defered in the function.")
	fmt.Println("Undefered in the function.")

}
