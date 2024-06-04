package main

import (
  "fmt"
  "time"
)

func main() {
	startTime := time.Now()
  answer := ""
  fmt.Println(answer)
	duration := time.Since(startTime)
	fmt.Println("Duration:", duration)
}
