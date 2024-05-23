// Starting in the top left corner of a 2 x 2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

// https://projecteuler.info/resources/images/0015.png

// How many such routes are there through a 20 x 20 grid?

package main

import "fmt"

func main() {
  answer := uint(0)
	const gridSize = 20
	paths := make([][]uint, gridSize+1)
	for i := range paths {
		paths[i] = make([]uint, gridSize+1)
	}

	for i := 0; i <= gridSize; i++ {
		paths[i][0] = 1
		paths[0][i] = 1
	}

	for i := 1; i <= gridSize; i++ {
		for j := 1; j <= gridSize; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	answer = paths[gridSize][gridSize]
  fmt.Println(answer)
	// 137846528820
}


// Solution using Combinatorics
// the number of moves will always be 40.  20 right and 20 down.  Can use combinations to get all possible paths.
// nCr = n! / (r! (n-r)!)
// Where n is the total number of items (40) and r is the number of items to choose (20 down moves in this case)
// nCr = 40! / (20! * (40-20)!) = 40! / (20! * 20!) = 137846528820

