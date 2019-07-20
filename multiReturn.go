package main

import "fmt"

func main() {
	n1, l1 := FullName("Zaikah", "Brownshed")
	fmt.Printf("Fullname: %v, Number of chars: %v%", n1, l1)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	lenght := len(full)
	return full, length
}
