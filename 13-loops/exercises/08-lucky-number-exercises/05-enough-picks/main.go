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
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

const (
	maxTurns = 5
	usage    = `This is a Lucky Number Game
[Enough Picks]

Put your guess number and will see if you will win the game in %d tries.

Let's Play!
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number. Try again.")
		return
	}

	if guess < 0 {
		fmt.Println("Positive numbers only.")
		return
	}

	gen := guess

	if gen < 10 {
		gen = 10
	}

	for i := 0; i < maxTurns; i++ {
		n := rand.Intn(gen + 1)
		fmt.Println(">", n)
		if n == guess {
			if i == 0 {
				fmt.Println("First Turn Winner: Lucky Guy!")
			} else {
				fmt.Println("You Won!")
			}
			return
		}
	}
	fmt.Printf("Sorry, you don't have luck.\nTry again!\n")
}
