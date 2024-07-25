package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// rand using math
	//fmt.Println(rand.Intn(5) + 1)

	// rand using crypto
	randNum, _ := rand.Int(rand.Reader, big.NewInt(5+1))
	fmt.Println(randNum)
}
