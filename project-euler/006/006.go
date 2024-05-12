package main

import "fmt"

func main() {
  answer := 0
  sumOfSquare := 0
  squareOfSum := 0
  for x := 1; x <= 100; x++ {
    sumOfSquare += x * x
    squareOfSum += x
  }
  squareOfSum *= squareOfSum
  fmt.Println(sumOfSquare)
  fmt.Println(squareOfSum)
  answer = squareOfSum - sumOfSquare
  fmt.Println(answer)
  // 25164150
}
