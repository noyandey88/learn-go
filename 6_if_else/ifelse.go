package main

import "fmt"

func main() {
	age := 10

	if age >= 18 {
		fmt.Println("You are adult")
	}else if age >= 12{
		fmt.Println("You are teenager")
	}	else{
		fmt.Println("You are a kid")
	}
	
}
