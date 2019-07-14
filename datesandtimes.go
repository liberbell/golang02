package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, time.June, 20, 11, 17, 0, 0, time.UTC)
	fmt.Printf("Juno was born: %s", t)
}
