package main

import "fmt"

func main() {
	// the defer keyword cuts whatever comes after it (in one line) and puts it just before the curly brace ends
	defer fmt.Println("This is defer :(")
	fmt.Println("Welcome to defer in go")
	deferTrigger()
}

func deferTrigger() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
