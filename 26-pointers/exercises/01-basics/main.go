// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var c *computer

	// compare the pointer variable to a nil value, and say it's nil
	if c == nil {
		fmt.Printf("c pointer is nil.\n")
	}

	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "appleee"}

	// put the apple into a new pointer variable
	newA := apple

	// compare the apples: if they are equal say so and print their addresses
	if newA == apple {
		fmt.Printf("Address are equal. newA[%p]: apple[%p]\n", newA, apple)
	}

	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{brand: "sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if apple != sony {
		fmt.Printf("they are not equal: apple[%p]::sony[%p]\n", apple, sony)
	}

	// put apple's value into a new ordinary variable
	appleVal := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple: add[%p], addp[%p]\n", &apple, apple)
	fmt.Printf("appleVal[%p]\n", &appleVal)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *apple == appleVal {
		fmt.Printf("Apple values are equal\n")
	}

	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Printf("*apple[%+v], newApple[%+v]\n", *apple, appleVal)

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sony, "lenovo")
	fmt.Printf("Sony's new brand is %s\n", sony.brand)

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	fmt.Printf("The value is: %+v\n", valueOf(apple))
	fmt.Printf("The value is: %+v\n", valueOf(sony))

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("New com: addr[%p]\n", newComputer("dell"))
	fmt.Printf("New com: addr[%p]\n", newComputer("acer"))
	fmt.Printf("New com: addr[%p]\n", newComputer("alienware"))
}

func change(old *computer, newBrand string) {
	old.brand = newBrand
}

func valueOf(com *computer) computer {
	return *com
}

func newComputer(new string) *computer {
	return &computer{brand: new}
}
