package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[2] = "Cherry"
	fruitList[3] = "Apple"

	fmt.Println(fruitList)
	fmt.Println("Length of the array", len(fruitList))

	var vegList = [3]string{"potato", "banana", "mushroom"}

	fmt.Println("Veggies List: ", vegList)
}
