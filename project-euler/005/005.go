package main

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

import "fmt"

func main() {
  answer := 20
  divisors := []int{20,19,18,17,16,15,14,13,12,11,7}
  for {
    isDivisible := true
    for i := range  divisors {
      if answer % divisors[i] != 0 {
        isDivisible = false
        break
      }
    }
    if isDivisible {
      break
    }

    answer++
  }
  fmt.Println(answer)
}

// Accidentally found prime factors instead of divisors.
// func main2() {
//   answer := 1
//   nums := make([]int, 20)

//   for x := range nums {
//     nums[x] = x + 1
//   }

//   nums = nums[1:]

//   fmt.Println(nums)

//   i := 0
//   for i < len(nums) {
//     for x := i + 1; x < len(nums); x++ {
//       if nums[x] % nums[i] == 0 {
//         nums = append(nums[:x], nums[x+1:]...)
//       }
//     }
//     i += 1
//   }

//   for x := range nums {
//     answer = answer * nums[x]
//   }

//   fmt.Println(answer)
// }