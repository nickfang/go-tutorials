package main

// By listing the first six prime numbers: 2, 3, 5, 7, 11, 13, we can see that the 6th prime is 13.
// What is the 10001st prime number?

import "fmt"

func main() {
  answer := 0
  primes := []int{2,3,5,7,11,13}
  num := 15
  for len(primes) < 10001 {
    isPrime := true
    for _, prime := range primes {
      if num % prime == 0 {
        isPrime = false
        break
      }
    }
    if isPrime {
      primes = append(primes, num)
      answer = num
    }
    // all even numbers are not prime
    num += 2
  }
  fmt.Println(answer)
  // 104743
}
