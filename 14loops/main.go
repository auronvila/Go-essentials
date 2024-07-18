package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//for index, day := range days {
	//	fmt.Printf("index is %v values is %v \n", index, day)
	//}

	rougeVal := 1
	for rougeVal < 10 {
		if rougeVal == 2 {
			goto lco
		}
		if rougeVal == 5 {
			break // breaks the loop we can use continue to skip over one step
		}
		fmt.Printf("Val is: %v \n", rougeVal)
		rougeVal++
	}

lco:
	fmt.Println("Jumping here")
}
