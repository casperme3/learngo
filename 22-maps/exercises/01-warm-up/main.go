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
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	phones := map[string]string{}
	phones["nolan"] = "11111"
	phones["saudia"] = "22222"
	phones["ylan"] = "33333"
	fmt.Printf("Phone directory: %q\n", phones)

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	products := map[int]bool{
		123: true,
		456: false,
		789: true,
		100: false,
	}
	fmt.Printf("Products available: %v\n", products)

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	nphones := map[string][]string{
		"nolan": {"111", "222", "333"},
		"yan2x": {"444", "555"},
		"ylan":  {"777"},
	}
	fmt.Printf("Phone numbers: %v\n", nphones)

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	shop := map[int]map[int]int{
		101: {90000: 5},
		102: {90011: 15},
		103: {90022: 25},
		104: {90033: 35},
	}
	fmt.Printf("Shopping basket: %v\n", shop)

}
