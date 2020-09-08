// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Refactor
//
//  Goal is refactoring the clock project by moving some of its parts to
//  a new file in the same package.
//
//  1. Create a new file: placeholders.go
//  2. Move the placeholder type to placeholder.go
//  3. Move all the placeholders (zero to nine and the colon) to placeholder.go
//  4. Move the digits array to placeholders.go
//
// HINT
//  + placeholders.go file should belong to main package as well
//
//    To remember how to do so: Rewatch the "PART I — Packages, Scopes and Importing"
//    section.
//
//  + Short declaration won't work in the package scope.
//    Remember why by watching: "Short Declaration: Package Scope" lecture
//
//  + If you receive the following error and you don't know what to do watch:
//    "Packages - Learn how to run multiple files" lecture.
//
//    # command-line-arguments
//    undefined: placeholder
//    undefined: colon
//
// EXPECTED OUTPUT
//  The same output before the refactoring.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const (
	sepa1 = iota + 2
	_
	_
	sepa2
)

func main() {
	// screen.Clear()
	//Infinite Loop: updates every 1 second
	for {
		screen.Clear()
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		// fmt.Println(now)

		clock := [...]blocktype{
			digits[hour/10], digits[hour%10],
			tickOn,
			digits[min/10], digits[min%10],
			tickOn,
			digits[sec/10], digits[sec%10],
		}

		if (sec % 2) == 0 {
			clock[sepa1], clock[sepa2] = tickOff, tickOff
		}

		for line := range clock[0] {
			for num := range clock {
				fmt.Print(clock[num][line], "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
