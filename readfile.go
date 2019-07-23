package main

func main() {
	filename := "/hello.txt"
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
