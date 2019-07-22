package main

import "fmt"

var isConnected bool = false

func main() {
	fmt.Printf("connection open: %v\n", isConnected)
	dosomething()
	fmt.Printf("connection open: %v\n", isConnected)
}

func dosomething() {
	connect()
	fmt.Println("defering connection.")
	defer disconnect()
	fmt.Printf("connection open: %v\n", isConnected)
	fmt.Println("Doing something.")
}

func connect() {
	isConnected = true
	fmt.Println("connected to database.")
}
