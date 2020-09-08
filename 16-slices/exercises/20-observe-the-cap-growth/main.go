// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the capacity growth
//
//  Write a program that appends elements to a slice
//  10 million times in a loop. Observe how the capacity of
//  the slice changes.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 10e6 times
//
//  3. On each iteration: Append an element to the slice
//
//  4. Print the length and capacity of the slice "only"
//     when its capacity changes.
//
//  BONUS: Print also the growth rate of the capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------
const max = 1e7

func main() {
	var (
		slice    []int
		capacity float64
	)
	// slice = append(slice, 1)
	// fmt.Printf("%d:%d\n", slice, cap(slice))

	for i := 0; i < max; i++ {
		newcap := float64(cap(slice))
		if newcap == 0 || capacity != newcap {
			fmt.Printf("len:%-15d cap:%-15d growth:%-5.2f \n", len(slice), cap(slice), newcap/capacity)
		}
		capacity = newcap
		slice = append(slice, i)
	}
	fmt.Printf("END\n")
}
