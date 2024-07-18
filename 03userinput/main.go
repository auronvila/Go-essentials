package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeText := "Welcome to user input"
	fmt.Println(welcomeText)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating for our Pizza: ")

	// comma ok || _ _
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
