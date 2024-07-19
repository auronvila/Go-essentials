package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go")

	content := "This needs to go in a file"
	file, err := os.Create("./mygoFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println(length)

	defer file.Close()

	readFile("./mygoFile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println(string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
