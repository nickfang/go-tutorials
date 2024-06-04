package utils

import (
	"sort"
)

func GetFactors(num int) []int {
	factors := []int{1, num}
	for i := 2; i*i < num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			if i != num/i {
				factors = append(factors, num/i)
			}
		}
	}
	sort.Ints(factors)
	return factors
}

func GetProperDivisors(num int) []int {
	factors := GetFactors(num)
	factors = factors[:len(factors)-1]
	return factors
}
