package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	//go greeter("Hello")
	//greeter("World")
	websiteList := []string{"https://google.com", "https://go.dev", "https://fb.com"}
	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
}

//func greeter(s string) {
//	for i := 0; i < 6; i++ {
//		fmt.Println(s)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()

	mut.Lock()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	mut.Unlock()
	fmt.Println("status code of the endpoint", res.StatusCode)
}
