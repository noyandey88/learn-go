package main

import "fmt"

const name string = "bangladesh"
const age = 26

// not allowed
// name := "golang"

func main() {
	// const name string = "bangladesh"
	// const age = 26
	// fmt.Println(age)
	const (
		port = 8000
		host = "localhost"
	)

	fmt.Println(port, host)
}