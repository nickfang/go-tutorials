// You are given the following information, but you may prefer to do some research for yourself.

// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
package main

import "fmt"

func GetDaysInMonth(month, year int) int {
	numDaysPerMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if month == 2 {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return numDaysPerMonth[2] + 1
		}
	}
	return numDaysPerMonth[month]
}

func main() {
	answer := 0

	// 0=sunday ... 6=saturday
	dayOfWeek := 1
	for year := 1900; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			if dayOfWeek == 0 && year >= 1901 {
				answer += 1
				fmt.Println(month, year)
			}
			daysInMonth := GetDaysInMonth(month, year)
			dayOfWeek = (dayOfWeek + daysInMonth) % 7
		}
	}
	fmt.Println(answer)
	// 171
}
