package main

import "fmt"

func main() {
	// var nums [4]int

	// nums[0] = 10
	// nums[1] = 20
	// nums[2] = 30
	// nums[3] = 40

	// fmt.Println(nums)

	names := [...]string{"javascript", "go", "java", "c#", "typescript"}

	fmt.Println(names)

	 cars := [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	 fmt.Println(cars)
	 
	
	// fmt.Println(len(nums))

	//  2d arrays
	nums := [2][2]int{{3, 4}, {5, 6}}

	fmt.Println(nums)
}