package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, time.June, 20, 11, 17, 0, 0, time.UTC)
	fmt.Printf("Juno was born: %s\n", t)

	now := time.Now()
	fmt.Printf("Now is: %s\n", now)

	fmt.Println("The month is: ", t.Month())
	fmt.Println("The day is: ", t.Day())
	fmt.Println("The weekday is: ", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Thursday, June 20, 2019"
	fmt.Println("Juno was born: ", t.Format(longFormat))
}
