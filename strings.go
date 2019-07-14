package main

import "fmt"

func main() {
	str1 := "An implicitly typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)

	str2 := "An explicitly typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)
}
