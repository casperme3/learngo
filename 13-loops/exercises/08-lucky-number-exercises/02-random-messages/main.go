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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

const (
	maxTurns = 5
	usage    = `This is a Lucky Number Game
[Random Messages]

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
		fmt.Println("Give positive numbers only.")
		return
	}

	for i := 0; i < maxTurns; i++ {
		n := rand.Intn(guess + 1)
		//fmt.Println(">", n)

		if n != guess {
			continue
		}

		switch rand.Intn(3) {
		case 0:
			fmt.Println("Great, you win!")
		case 1:
			fmt.Println("Awesome, you are lucky!")
		case 2:
			fmt.Println("You're are born a winner!")
		}
		return

	}

	switch rand.Intn(3) {
	case 0:
		fmt.Println("No chance. Try again!")
	case 1:
		fmt.Println("So sad you lost!")
	case 2:
		fmt.Println("You're a natural loser!")
	}
}
