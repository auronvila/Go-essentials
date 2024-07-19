package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Lco web request")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response", res)
	defer res.Body.Close()

	dataBytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes))
}
