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
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	splitA := strings.Split(namesA, ", ")

	fmt.Printf("%#v\n", splitA)
	fmt.Printf("%#v\n", namesB)
	sort.Strings(splitA)
	sort.Strings(namesB)

	if len(splitA) == len(namesB) {
		for i, v := range splitA {
			if v != namesB[i] {
				fmt.Printf("Slices are not equal.\n")
				return
			}
		}
		fmt.Printf("Yes, slices are equal.\n")
	}
}
