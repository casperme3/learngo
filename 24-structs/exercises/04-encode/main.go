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
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

func main() {
	type exported struct {
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
		case "save":
			////////////////////////////
			//NOOB STYLE:
			// it, saved := 0, make([]exported, len(games))
			// for _, v := range games {
			// 	saved[it].ID = v.id
			// 	saved[it].Name = v.name
			// 	saved[it].Genre = v.genre
			// 	saved[it].Price = v.price
			// 	it++
			// }

			//EXPERT STYLE:
			var saved []exported
			for _, v := range games {
				saved = append(saved, exported{v.id, v.name, v.genre, v.price})
			}
			////////////////////////////

			out, err := json.MarshalIndent(saved, "", "\t")
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(out))
			return
		}
	}

}
