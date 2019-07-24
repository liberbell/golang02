package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	resp, err := http.Get(url)
	checkError(err)
}
