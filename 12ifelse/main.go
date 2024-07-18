package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 32

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Unhandled msg"
	}

	fmt.Println(result)
}
