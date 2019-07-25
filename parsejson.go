package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	// fmt.Println(content)
	tours := tours.json(content)
	fmt.Println(tours)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	jsondata := string(bytes)
	fmt.Print(jsondata)
}

func tourFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkerror(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkerror(err)
		tours = append(tours, tour)
	}
	return tours
}
