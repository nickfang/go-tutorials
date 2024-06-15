package main

// The prime factors of 13195 are 5, 7, 13 and 29
// What is the largest prime factor of the number 600851475143

import (
	"fmt"
	"math"
)

func main() {
	largestFactor := 1
	number := 600851475143

	if number%2 == 0 {
		number = (number / 2)
		largestFactor = 2
	}

	end := int(math.Sqrt(float64(number)))

	for x := 3; x < end; x += 2 {
		if number%x == 0 {
			number = number / x
			largestFactor = x
		}
	}
	fmt.Println(largestFactor)
	// 6857
}
