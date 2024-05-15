package main

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, a^2 + b^2 = c^2.

// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

import (
	"fmt"
	"math"
)

func main() {
  answer := 0
  a := 0
  b := 0
  c := float64(0)
  loops := 0
  Outer:
    // solved for a^2+b^2=c^2 and a+b+c=1000 and substituted 1 for b to get upper limit for loops.
    for a = 2; a < 500; a++ {
      for b = 2; b < 500; b++ {
        cSquared := (a*a) + (b*b)
        c = math.Sqrt(float64(cSquared))

        loops++
        if a + b + int(c) > 1000 {
          // optimization: does 87865 loops as opposed 98978 loops without check.
          break
        }

        if c == math.Floor(c) {
          fmt.Println(a, b, c)
          if a + b + int(c) == 1000 {
            answer = a * b * int(math.Floor(c))
            break Outer
          }
        }
      }
    }
  fmt.Println(answer)
  // 31875000
}
