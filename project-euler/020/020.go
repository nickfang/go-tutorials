// n! means n X (n - 1) x ... x 3 x 2 x 1.

// For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

// Find the sum of the digits in the number 100!.

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	answer := 0

	// calculate factorial
	lastFactorial := big.NewInt(1)
	factorial := big.NewInt(1)
	for i := 1; i <= 100; i++ {
		num := big.NewInt(int64(i))
		factorial.Mul(num, lastFactorial)
		lastFactorial = factorial
	}

	// sum the digita
	for _, digit := range factorial.String() {
		num, _ := strconv.Atoi(string(digit))
		answer += num
	}

	fmt.Println(answer)
	// 648
}
