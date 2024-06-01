// If the numbers 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundredand fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

package main

import "fmt"

func getNumString(num int) string {
	ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	teens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	numString := ""
	if num < 10 {
		numString = ones[num]
	} else if num < 20 {
		numString = teens[num-10]
	} else if num < 100 {
		numString = tens[num/10] + getNumString(num%10)
	} else if num < 1000 {
		numString = ones[num/100] + "hundred" + (func() string {
			if num%100 == 0 {
				return ""
			} else {
				return "and" + getNumString(num%100)
			}
		}())
	} else {
		numString = "onethousand"
	}
	return numString
}

func main() {
	answer := 0
	for i := 0; i <= 1000; i++ {
		numString := getNumString(i)
		answer += len(numString)
	}

	fmt.Println(answer)
	// 21124
}
