package main

import "fmt"

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)
	// if err == nil {
	fmt.Println("string length:", stringLength)
	// }

	fmt.Printf("value of aNumber: %v\n", aNumber)
	fmt.Printf("value of isTrue: %v\n", isTrue)
	fmt.Printf("value of aNumber as float: %.2f\n", float64(aNumber))

	fmt.Printf("Data types: %T, %T, %T, %T and %T\n", str1, str2, str3, isTrue, aNumber)

	myString := fmt.Sprintf("Data types: %T, %T, %T, %T and %T\n", str1, str2, str3, isTrue, aNumber)

}
