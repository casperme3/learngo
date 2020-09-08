// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	one, two := 1.11, 2.22
	swapVal(&one, &two)
	fmt.Printf("one:%g, two:%g\n", one, two)

	p1, p2 := &one, &two
	// p1, p2 = p2, p1
	fmt.Printf(">> Bef: p1[%p], p2[%p]\n", p1, p2)
	// swapAddr(p1, p2)
	p1, p2 = swapAddr2(p1, p2)
	// swapAddr(&p1, &p2) //for double pointer
	fmt.Printf("======================\n")
	fmt.Printf("Addr: one[%g], two[%g]\n", one, two)
	fmt.Printf("Addr: p1[%p], p2[%p]\n", p1, p2)
	fmt.Printf("Value: p1[%g], p2[%g]\n", *p1, *p2)
}

func swapVal(var1, var2 *float64) {
	*var1, *var2 = *var2, *var1
}

//will only swap the value, not the address
func swapVal2(var1, var2 *float64) {
	// fmt.Printf(">> Aft: v1[%p], v2[%p]\n", var1, var2)
	*var1, *var2 = *var2, *var1
}

//Double pointer is working... WOW!
//This will swap the Address pointed by both pointer
func swapAddr(var1, var2 **float64) {
	*var1, *var2 = *var2, *var1
}

func swapAddr2(var1, var2 *float64) (*float64, *float64) {
	return var2, var1
}
