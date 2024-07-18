package main

import "fmt"

// ArrSize -> the first letter is capital it means that the variable is public
const ArrSize int = 10

func main() {
	var username string = "auronvila"
	fmt.Printf(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	print(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 100000
	print(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 1.2
	print(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and aliases
	var anotherVar float32
	print(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	// implicit type
	var implicitVar = "I am initialized implicitly"
	fmt.Println(implicitVar)

	// no var style
	numberOfUsers := 3000
	fmt.Printf("This is an int %D", numberOfUsers)
}
