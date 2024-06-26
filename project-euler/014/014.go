// The following iterative sequence is defined for the set of positive integers:
//   n -> n/2 (n is even)
//   n -> 3n+1 (n is odd)
// Using the rule above and starting with 13, we generate the following sequence:
//   13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
// Which starting number, under one million, produces the longest chain?
//NOTE: Once the chain starts the terms are allowed to go above one million.

package main

import "fmt"

func main() {
  answer := 0
  maxLength := 0
  for i := 999999; i > 13; i-- {
    chainLength := 1
    x := i
    for x > 1 {
      if x % 2 == 0 {
        x = x / 2
      } else {
        x = 3 * x + 1
      }
      chainLength += 1
    }
    if chainLength > maxLength {
      maxLength = chainLength
      answer = i
    }
  }

  fmt.Println(answer)
  // 837799
}
