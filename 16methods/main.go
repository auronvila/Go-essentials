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
	auron.GetStatus()
	auron.NewMail()
	fmt.Printf("Aurons email is %v", auron.Email)
}

func (u User) GetStatus() {
	fmt.Println("Is user active: \n", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
