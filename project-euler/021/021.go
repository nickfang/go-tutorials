// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).<br>
// If d(a) = b and d(b) = a, where a not equal b, then a and b are an amicable pair and each of a and b are called amicable numbers.
// For example:
//   the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
// therefore d(220) = 284.
//   The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
// Evaluate the sum of all the amicable numbers under 10000.

package main

import (
	"fmt"
	"project-euler-go/utils"
	// "slices"
	"time"
)

func SumDivisors(num int) int {
	factors := utils.GetProperDivisors(num)
	sum := 0
	for _, factor := range factors {
		sum += factor
	}
	return sum
}

func main() {
	startTime := time.Now()
	answer := 0
	allNums := make([]int, 10000)
	amicableNums := []int{}
	for x := range allNums {
		allNums[x] = -1
	}
	for x, num := range allNums {
		if num != -1 {
			continue
		}
		sum := SumDivisors(x)
		sum2 := SumDivisors(sum)
		if x == sum2 && x != sum {
			amicableNums = append(amicableNums, x)
			amicableNums = append(amicableNums, sum)
			allNums[x] = sum
			allNums[sum] = x
		}
	}
	for _, num := range amicableNums {
		answer += num
	}
	fmt.Println(answer)
	duration := time.Since(startTime)
	fmt.Println("Duration:", duration)
	// 31626
}
