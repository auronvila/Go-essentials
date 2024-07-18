package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and case in Go")

	randomNum := rand.Intn(6) + 1

	fmt.Printf("Val of dice is %v \n", randomNum)
	switch randomNum {
	case 1:
		fmt.Println("The value is 1")
	case 2:
		fmt.Println("The value is 2")
	case 3:
		fmt.Println("The value is 3")
		fallthrough
	case 4:
		fmt.Println("The value is 4")
		// the fallthrough will run the next case if it is inside of this case.
		fallthrough
	case 5:
		fmt.Println("The value is 5")
	case 6:
		fmt.Println("The value is 6")
	default:
		fmt.Printf("Unhandled number %v \n", randomNum)
	}
}
