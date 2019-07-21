package main

type Animal interface {
  Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
  retrun "woof!"
}
