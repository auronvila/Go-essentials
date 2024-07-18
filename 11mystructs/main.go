package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in Go")
	auron := User{
		Name:   "Auron",
		Email:  "auron@gmail.com",
		Status: true,
		Age:    19,
	}
	fmt.Printf("Auron %v", auron)
	// print with details
	fmt.Printf("Auron details %+v \n", auron)
	// print single val
	fmt.Printf("Auron name %v \n", auron.Name)
}
