// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------
type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	games := []game{
		{item: item{id: 1, name: "pubg", price: 20}, genre: "action"}, //complete with types
		{item: item{3, "com2", 10}, genre: "stat"},                    //inner struct dont use types
		{item{5, "ML", 20}, "moba"},                                   //shortcut no types at all
		{item{6, "roblox", 5}, "sandbox"},
	}

	in := bufio.NewScanner(os.Stdin)
	// in.Split(bufio.ScanWords)

	fmt.Printf("Nolan's store have [%d] games.\n", len(games))

	for {
		fmt.Printf("\n  > list  : to list all games\n")
		fmt.Printf("  > quit  : to quit program\n\n")

		in.Scan()

		switch cmd := in.Text(); cmd {
		case "list":
			for _, v := range games {
				fmt.Printf("#%-4d %-10q %5d    %-15s\n", v.id, v.name, v.price, v.genre)
			}
		case "quit":
			fmt.Println("\nProgram exited!")
			return
		}
	}
}
