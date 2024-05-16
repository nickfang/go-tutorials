// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.
package main

import "fmt"

func main() {
  answer := 0
  primes := []int{2,3,5,7,11,13}
  num := 15
  for num < 2000000 {
    isPrime := true
    for _, prime := range primes {
      if num % prime == 0 {
        isPrime = false
        break
      }
    }
    if isPrime {
      primes = append(primes, num)
    }
    // all even numbers are not prime
    num += 2
  }
  fmt.Println(primes)
  for _, prime := range primes {
    answer += prime
  }
  fmt.Println(answer)
}

// Alternative method that's much faster
// func main() {
//   answer := 0
//   max := 2000000
//   sieve := make([]bool, max + 1)
//   for x := 2; x <= max; x++ {
//     sieve[x] = true
//   }
//   for x := 2; x*x <= max; x++ {
//     if sieve[x] {
//       for y := x*x; y <= max; y += x {
//         sieve[y] = false
//       }
//     }
//   }
//   for x := 2; x <= max; x++ {
//     if sieve[x] {
//       answer += x
//     }
//   }
//   fmt.Println(answer)
// }
