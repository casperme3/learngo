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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------
const (
	maxTurns = 5
	usage    = `This is a Lucky Number Game
[Double Guesses]

Put your guess (2) numbers and will see if you will win the game in %d tries.

Let's Play!
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	val1, err1 := strconv.Atoi(args[0])
	val2, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Put a valid number.")
		return
	}

	if val1 < 0 || val2 < 0 {
		fmt.Println("Give positive numbers only.")
		return
	}

	guess := val1
	if val2 > val1 {
		guess = val2
	}
	// fmt.Println("guess: ", guess)
	for i := 0; i < maxTurns; i++ {
		n := rand.Intn(guess + 1)
		// fmt.Println(">", n)
		if n == val1 && n == val2 {
			fmt.Printf("Great you win [%d]\n", n)
			return
		}
	}
	fmt.Println("Sorry you lost, even with 2 numbers.")
}
