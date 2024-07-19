package main

import "fmt"

func test() string {
	str := "Auron"
	return str
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	returnedData := test()
	summedData := sum(2, 5)
	proData := proAdder(1, 2, 3, 4, 5, 6)
	fmt.Println(returnedData)
	fmt.Println(summedData)
	fmt.Println(proData)
}
