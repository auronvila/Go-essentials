package main

import (
	"fmt"
	"github.com/auronvila/router"
	"log"
	"net/http"
)

func main() {
	const port string = ":4000"
	fmt.Println("MongoDB API")

	err := http.ListenAndServe(port, router.Router())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening at port: ", port)
}
