package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{}
	fruitList = append(fruitList, "Test", "Test1", "Test2", "Test3")
	fruitList = append(fruitList[:3])
	//fmt.Println(fruitList)

	// different way of declaring an array
	var highScores = make([]int, 4)
	var i int = 0
	for i < 4 {
		highScores[i] = i
		i++
	}

	highScores = append(highScores, 4, 6, 5, 10, 9)
	sort.Ints(highScores)

	//fmt.Println(highScores)

	// removing an item from the arr
	var courseList = []string{"react", "ts", "js", "go", "C"}
	var index int = 2
	courseList = append(courseList[:index], courseList[index+1:]...)
	fmt.Println(courseList)
}
