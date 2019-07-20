package main

import "fmt"

func main() {
	n1, l1 := FullName("Zaikah", "Brownshed")
	fmt.Printf("Fullname: %v, Number of chars: %v\n", n1, l1)

	n1, l1 := FullNameNakedReturn("Zaikah", "Brownshed")
	fmt.Printf("Naked Fullname: %v, Number of chars: %v\n", n1, l1)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f, l string) (full string, length int) {
	full := f + " " + l
	length := len(full)
	return
}
