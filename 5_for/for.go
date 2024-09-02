package main

import "fmt"

// for -> only construct in go for lopping

func main(){
	// while loop
	i := 1
	for i <= 6{
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for{
	// 	println("1")
	// }

	fmt.Println("----------------------------------------------------------------")

	// for loop
	for i := 0; i <= 3; i++{
		// break;

		if i == 2{
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----------------------------------------------------------------")

	// 1.22 range
	for i:= range 3{
		fmt.Println(i)
	}
}