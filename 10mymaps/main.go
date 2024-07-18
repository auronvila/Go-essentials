package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Go")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println(languages)
	fmt.Println("js: ", languages["js"])

	// delete a record in map
	//delete(languages, "js")
	//fmt.Println(languages)

	for key, val := range languages {
		fmt.Printf("Key: %v, Val: %v \n", key, val)
	}
}
