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
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file with their ordinals
//
//  Use the previous exercise.
//
//  This time, print the sorted items with ordinals
//  (see the expected output)
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     1. apple
//     2. banana
//     3. orange
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + You can use strconv.AppendInt function to append an int
//     to a byte slice. strconv contains a lot of functions for appending
//     other basic types to []byte slices as well.
//
//   + You can append individual characters to a byte slice using
//     rune literals (because: rune literal are typeless numerics):
//
//     var slice []byte
//     slice = append(slice, 'h', 'i', ' ', '!')
//     fmt.Printf("%s\n", slice)
//
//     Above code prints: hi !
// ---------------------------------------------------------
func main() {
	args := os.Args[1:]
	s.PrintBacking = true

	if len(args) == 0 {
		fmt.Println("Give me some items to sort.")
		return
	}
	sort.Strings(args)

	var total int
	for _, arg := range args {
		total += len(arg) + 4 //4 means the ordinal place and the newline
	}
	fmt.Printf("total: %d\n", total)

	sorted := make([]byte, 0, total)
	s.Show("init: ", sorted)

	for i, arg := range args {
		sorted = strconv.AppendInt(sorted, int64(i+1), 10)
		sorted = append(sorted, '.', ' ')
		sorted = append(sorted, arg...)
		sorted = append(sorted, '\n')
	}

	//TODO:
	// why cant mix this: append(sorted, '.', ' ', arg...)
	// Also check the address of the backing array before and after if it changes

	err := ioutil.WriteFile("sorted2.txt", sorted, 0644)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	s.Show("end: ", sorted)
}

func init() {
	s.MaxPerLine = 10
	s.Width = 45
}
