package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speark() {
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 37, "woof"}
	fmt.Println(poodle)
	poodle.Speak()
}
