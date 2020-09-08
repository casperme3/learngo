// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

var atty int = 3

func main() {
	var attyw int = 2

	_ = attyw

	atty := 5
	fmt.Println(atty)

	atty += 3
	fmt.Println(atty)

	next()
	fmt.Println("aft next:", atty)
}

func next() {
	fmt.Println("next:", atty)
	atty += 2
}
