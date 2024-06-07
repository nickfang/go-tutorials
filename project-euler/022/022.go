// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

// What is the total of all the name scores in the file?

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	answer := 0

	data, err := os.ReadFile("022/names.txt")
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("022.go should be run from the project-euler directory.")
		return
	}

	names := strings.Split(string(data), ",")
	for i := range names {
		// Remove white space and quotes
		names[i] = strings.Trim(strings.TrimSpace(names[i]), "\"")
	}

	sort.Strings(names)

	for i, name := range names {
		nameScore := i + 1
		nameValue := 0
		for _, char := range name {
			nameValue += int(char - 'A' + 1)
		}
		score := nameValue * nameScore
		answer += score
		fmt.Println(name, score)
	}

	fmt.Println(answer)
	duration := time.Since(startTime)
	fmt.Println("Duration:", duration)
	// 871198282
}
