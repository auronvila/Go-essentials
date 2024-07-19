package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://google.com/learn?course=react&paymentid=gsabse334"

func main() {
	fmt.Println("Welcome to handling urls in go")
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)
	//https
	fmt.Println(result.Scheme)
	//learn
	fmt.Println(result.Path)
	//google.com
	fmt.Println(result.Host)
	//course=react&paymentid=gsabse334
	fmt.Println(result.RawQuery)

	params := result.Query()

	fmt.Println(params["course"])

	for _, val := range params {
		fmt.Println("Val of the param: ", val)
	}
}
