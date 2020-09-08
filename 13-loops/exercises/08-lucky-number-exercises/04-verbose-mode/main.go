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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	maxTurns = 5
	usage    = `This is a Lucky Number Game
[Verbose Mode]

Put your guess number and option mode, and will see if you will win the game in %d tries.
[mode][guess number]
Ex: -v 5

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

	guess, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Put a valid number.")
		return
	}

	if guess < 0 {
		fmt.Println("Give positive numbers only.")
		return
	}

	var verb bool
	if args[0] == "-v" {
		verb = true
	}

	// fmt.Println("guess: ", guess)
	for i := 0; i < maxTurns; i++ {
		n := rand.Intn(guess + 1)

		if verb {
			fmt.Printf("%d ", n)
		}
		if n == guess {
			fmt.Printf(":: Great you win [%d]\n", n)
			return
		}
	}
	fmt.Println(":: Sorry you lost, again!")
}
