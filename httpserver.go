package main

import "net/http"

type hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request)

func main() {
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
