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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "gods of wars",
                "genre": "action adventure",
                "price": 51
        },
        {
                "id": 2,
                "name": "xxx-com 3",
                "genre": "strategic",
                "price": 42
        },
        {
                "id": 3,
                "name": "minedcrack",
                "genre": "sandboxed",
                "price": 23        }
]`

func main() {
	type jsonGame struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
	}

	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	// games := []game{
	// 	{item: item{id: 1, name: "pubg", price: 20}, genre: "action"}, //complete with types
	// 	{item: item{3, "com2", 10}, genre: "stat"},                    //inner struct dont use types
	// 	{item{5, "ML", 20}, "moba"},                                   //shortcut no types at all
	// 	{item{6, "roblox", 5}, "sandbox"},
	// }

	in := bufio.NewScanner(os.Stdin)

	//Initialize games store from the constants
	var games []game
	var impt []jsonGame

	if err := json.Unmarshal([]byte(data), &impt); err != nil {
		fmt.Println("Error in  Unmarshal")
		return
	}

	//Store imported data to games slice
	for _, g := range impt {
		games = append(games, game{item{g.ID, g.Name, g.Price}, g.Genre})
	}

	fmt.Printf("Nolan's store have [%d] games.\n", len(games))

	for {
		fmt.Printf("\n  > list  : to list all games\n")
		fmt.Printf("  > id n  : to search for specific id\n")
		fmt.Printf("  > save  : to export data in json format\n")
		fmt.Printf("  > quit  : to quit program\n\n")

		if !in.Scan() {
			break
		}

	outer:
		switch cmd := strings.Fields(in.Text()); cmd[0] {
		case "quit":
			fmt.Println("\nProgram exited!")
			return
		case "list":
			for _, v := range games {
				fmt.Printf("#%-4d %-15q %5d    %-25s\n", v.id, v.name, v.price, v.genre)
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
		case "save":
			var saved []jsonGame
			for _, v := range games {
				saved = append(saved, jsonGame{v.id, v.name, v.genre, v.price})
			}
			////////////////////////////

			expt, err := json.MarshalIndent(saved, "", "\t")
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(expt))
			return
		}
	}
}
