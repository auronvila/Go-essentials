package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers class")

	myNum := 23
	var ptr *int = &myNum

	fmt.Println(*ptr, ptr)

	*ptr += 4
	fmt.Println("Updated val: ", myNum)
}
