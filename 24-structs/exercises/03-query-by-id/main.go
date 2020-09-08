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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. i don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
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
		fmt.Printf("  > id n  : to search for specific id\n")
		fmt.Printf("  > quit  : to quit program\n\n")

		if !in.Scan() {
			break
		}

	outer:
		switch cmd := strings.Fields(in.Text()); cmd[0] {
		case "list":
			for _, v := range games {
				fmt.Printf("#%-4d %-10q %5d    %-15s\n", v.id, v.name, v.price, v.genre)
			}
		case "id":
			if len(cmd) == 2 {
				if n, err := strconv.Atoi(cmd[1]); err == nil {
					for _, v := range games {
						if v.id == n {
							fmt.Printf("#%-4d %-10q %5d    %-15s\n", v.id, v.name, v.price, v.genre)
							break outer
						}
					}
					fmt.Println("Sorry, id game not found!")
					break outer
				}
			}
			fmt.Println("wrong id detected. try again...")
		case "quit":
			fmt.Println("\nProgram exited!")
			return
		}
	}
}
