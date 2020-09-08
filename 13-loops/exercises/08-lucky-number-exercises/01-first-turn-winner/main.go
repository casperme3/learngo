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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------
const (
	maxTurns = 5
	usage    = `This is a Lucky Number Game
[First Turn winner]

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
		if n == guess {
			if i == 0 {
				fmt.Println("First Turn Winner: Lucky Guy!")
			} else {
				fmt.Println("You Won!")
			}
			return
		}
	}
	fmt.Printf("Sorry, you don't have what it takes.\nTry again!\n")
}
