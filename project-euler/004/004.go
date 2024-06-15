package main

// A palindromic number reads the same both ways.  The largest palindrome made from the product of 2-digit numbers is 9009 = 91 x 99
// Find the largest palindrome made from the product of two 3 digit numbers.

import "fmt"

func isPalindrome(num int) bool {
	numStr := fmt.Sprint(num)
	numStrLen := len(numStr)
	for i := 0; i < numStrLen/2; i++ {
		// don't need to check middle number if odd number of digits.
		if numStr[i] != numStr[numStrLen-1-i] {
			return false
		}
	}
	return true
}

func main() {
	answer := []int{0, 0}
	maxPalindrome := 0
	product := 0
	for x := 900; x < 1000; x++ {
		for y := 900; y < 1000; y++ {
			product = x * y
			if isPalindrome(product) {
				if product > maxPalindrome {
					maxPalindrome = product
					answer[0] = x
					answer[1] = y
				}
			}
		}
	}
	fmt.Println(maxPalindrome)
	fmt.Println(answer)
	// 906609
}
