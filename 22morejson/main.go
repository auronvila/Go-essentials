package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs Bootcamp", 299, "google.com", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "google.com", "bvf122", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 42, "google.com", "bngh323", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", finalJson)
}
