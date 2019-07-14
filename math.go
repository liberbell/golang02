package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	intsum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intsum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatsum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatsum)
}
