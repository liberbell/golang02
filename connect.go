package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Printf("connection open: %v\n", isConnected)
	dosomething()
	fmt.Printf("connection open: %v\n", isConnected)
}
