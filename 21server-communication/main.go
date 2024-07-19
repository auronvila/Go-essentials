package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseUrl string = "http://localhost:3100"

func main() {
	fmt.Println("Welcome to a 21server-communication written in GO")
	HandleGetReq()
	HandlePostReq()
}

func HandleGetReq() {
	res, _ := http.Get(baseUrl)

	responseString := strings.Builder{}
	defer res.Body.Close()
	// content is in byte format, so we need something to translate the bytecode
	//content, _ := io.ReadAll(res.Body)
	fmt.Println("Res body printed into using builder", responseString.String())

	//fmt.Println("Res status: ", res.Status)
	//fmt.Println("Res body: ", string(content))
}

func HandlePostReq() {
	reqBody := strings.NewReader(`
{
"name":"Auron"
}
`)
	res, _ := http.Post(baseUrl, "application/json", reqBody)
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
