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
	fmt.Println(num)
	for _, digitStr := range num.String() {
		fmt.Println(digitStr)
		asciiNum := digitStr - '0'
		fmt.Println(asciiNum)
		answer += int(asciiNum)
	}
  fmt.Println(answer)
}
