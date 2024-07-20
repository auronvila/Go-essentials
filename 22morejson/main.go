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
	//EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsoData := []byte(`
        {
                "courseName": "Mern Bootcamp",
                "price": 199,
                "platform": "google.com",
                "tags": ["full-stack","js"]
        }
`)

	//var Course course
	checkJson := json.Valid(jsoData)

	//if checkJson {
	//	fmt.Println("Json was valid")
	//	err := json.Unmarshal(jsoData, &Course)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Printf("%#v\n", Course)
	//} else {
	//	fmt.Println("Json was not valid")
	//}

	if checkJson {
		var onlideData map[string]interface{}
		err := json.Unmarshal(jsoData, &onlideData)
		if err != nil {
			panic(err)
		}
		for k, v := range onlideData {
			fmt.Printf("Key: %v, Value: %v Value type: %T \n", k, v, v)
		}

		//fmt.Printf("%#v\n", Course)
	} else {
		fmt.Println("Json was not valid")
	}
}
