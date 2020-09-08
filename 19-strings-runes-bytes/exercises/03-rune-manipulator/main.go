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
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	// Print the byte and rune length of the words
	// Hint: Use len and utf8.RuneCountInString

	// Print the bytes of the words in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the words in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the words as rune literals
	// Hint: Use for range

	// Print the first rune and its byte size of the words
	// Hint: Use utf8.DecodeRuneInString

	// Print the last rune of the words
	// Hint: Use utf8.DecodeLastRuneInString

	// Slice and print the first two runes of the words

	// Slice and print the last two runes of the words

	// Convert the string to []rune
	// Print the first and last two runes

	for _, str := range words {
		strlen := len(str)
		fmt.Printf("%q\n", str)
		fmt.Printf("\tit has %d bytes and %d runes.\n", strlen, utf8.RuneCountInString(str))
		fmt.Printf("\tbytes     : % x\n", []byte(str))
		fmt.Printf("\trunes     :")
		for _, r := range str {
			fmt.Printf("% x", r)
		}
		fmt.Println()
		fmt.Printf("\trunesl    :")
		for _, r := range str {
			fmt.Printf("%q", r)
		}
		fmt.Println()
		r, s := utf8.DecodeRuneInString(str)
		fmt.Printf("\tfirst     : %q (%d bytes)\n", r, s)

		r, s = utf8.DecodeLastRuneInString(str)
		fmt.Printf("\tlast      : %q (%d bytes)\n", r, s)

		_, f1 := utf8.DecodeRuneInString(str)
		_, f2 := utf8.DecodeRuneInString(str[f1:])
		fmt.Printf("\tfirst2    : %q\n", str[:(f1+f2)])

		_, l1 := utf8.DecodeLastRuneInString(str)
		_, l2 := utf8.DecodeLastRuneInString(str[:strlen-l1])
		fmt.Printf("\tlast2     : %q\n", str[strlen-l1-l2:])

		fmt.Printf("\t%s\n", strings.Repeat("-", 30))
		runes := []rune(str)
		fmt.Printf("\tfirst2    : %q\n", string(runes[:2]))
		fmt.Printf("\tlast2     : %q\n", string(runes[len(runes)-2:]))
	}
}
