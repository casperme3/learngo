// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got request logs of a web server. The log data
//  contains 8-hourly totals per each day. It is stored
//  in the `reqs` slice.
//
//  Find and print the total requests per day, as well as
//  the grand total.
//
//  See the `reqs` slice and the steps in the code below.
//
//
// RESTRICTIONS
//
//   1. You need to produce the daily slice, don't just loop
//      and print the element totals directly. The goal is
//      gaining more experience in slice operations.
//
//   2. Your code should work even if you add to or remove the
//      existing elements from the `reqs` slice.
//
//      For example, after solving the exercise, try it with
//      this new data:
//
//      reqs := []int{
// 	      500, 600, 250,
// 	      200, 400, 50,
// 	      900, 800, 600,
// 	      750, 250, 100,
// 	      150, 654, 235,
// 	      320, 534, 765,
// 	      121, 876, 285,
// 	      543, 642,
// 	      // the last element is missing (your code should be able to handle this)
// 	      // that is why you shouldn't calculate the `size` below manually.
//      }
//
//      The grand total of the new data should be 10525.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, // 4th day: 1100 requests
		// grand total: 5400 requests
	}

	// reqs := []int{
	// 	500, 600, 250, //1350
	// 	200, 400, 50, //650
	// 	900, 800, 600, //2300
	// 	750, 250, 100, //1100
	// 	150, 654, 235, //1039
	// 	320, 534, 765, //1619
	// 	121, 876, 285, //1282
	// 	543, 642, //1185
	// }

	// ================================================
	// #1: Make a new slice with the exact size needed.
	reql := len(reqs)
	size := int(math.Ceil(float64(reql) / 3.))
	// daily := make([][]int, size) //For 1st and 2nd solution
	daily := make([][]int, 0, size)
	// fmt.Printf("len:%d, size:%d\n", reql, size)
	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	///FIRST SOLUTION:
	// for i := 0; i < size; i++ {
	// 	daily[i] = make([]int, 0, N)
	// 	for n := 0; n < N; n++ {
	// 		if idx := i*N + n; idx < reql {
	// 			// fmt.Printf("%d, ", i*N+n)
	// 			daily[i] = append(daily[i], reqs[i*N+n])
	// 		}
	// 	}
	// }

	///SECOND SOLUTION:
	// for i := 0; i < size; i++ {
	// 	daily[i] = make([]int, 0, N)
	// 	stop := N
	// 	if len(reqs) < N {
	// 		stop = len(reqs)
	// 	}
	// 	daily[i] = append(daily[i], reqs[:stop]...)
	// 	reqs = reqs[stop:]
	// }

	///THIRD SOLUTION: almost same as Inanc's solution
	for i := 0; i < size; i++ {
		stop := N
		if len(reqs) < N {
			stop = len(reqs)
		}
		daily = append(daily, reqs[:stop])
		reqs = reqs[stop:]
	}

	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	var grand int
	for i, day := range daily {
		var total int
		for _, per := range day {
			total += per
			fmt.Printf("%-10d%-10d\n", i+1, per)
		}
		fmt.Printf("  Total: %d\n\n", total)

		grand += total
	}
	fmt.Printf("   GRAND: %d", grand)
	// ================================================
}

/*
1         500
1         600
1         250
   TOTAL: 1350

2         200
2         400
2         50
   TOTAL: 650
*/
