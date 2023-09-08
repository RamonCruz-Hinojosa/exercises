package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of x is:", i)
	// }

	names := []string{"mario", "luigi", "yoshi", "peach"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}

	// the value in the loop above doesnt affect the names object on line 18 it uses a copy in the loop
	fmt.Println(names)
}
