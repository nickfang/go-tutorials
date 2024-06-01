package main

import (
	"fmt"
	"math/big"
)

func main() {
	answer := 0
	// fmt.Println(2 << 14)  // playing with bit shifting
	num := new(big.Int)
	num.Exp(big.NewInt(2), big.NewInt(1000), nil) // 2^1000
	for _, digitStr := range num.String() {
		asciiNum := digitStr - '0'
		answer += int(asciiNum)
	}
	fmt.Println(answer)
	// 1366
}
