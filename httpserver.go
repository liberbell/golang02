package main

type hello struct {
}

func main() {
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
