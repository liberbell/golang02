package main

import "fmt"

func main() {
	var p *int

	if p != nil {
		fmt.Println("value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}
	var v int = 42
	p = &v
	if p != nil {
		fmt.Println("value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}
}
