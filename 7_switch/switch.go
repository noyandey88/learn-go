package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	i := 1

	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unknown")
	}

	// multiple condition switch

	fmt.Println(time.Now().Weekday())

	switch time.Now().Weekday(){
	case time.Friday, time.Saturday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's Workday")
	}

	// type switch
	whoAmI:= func(i interface{}){
		// switch t:= i.(type){
		switch i.(type){
		case int:
			fmt.Println("It is an integer")
		case string:
			fmt.Println("It is a string")
		case bool:
			fmt.Println("It is a boolean")
		case float32, float64:
				fmt.Println("It is a float")
		default:
			fmt.Println("It's an unknown type")
		}
	}
	
	whoAmI("golang")
}