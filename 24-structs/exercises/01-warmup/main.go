// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
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

	fmt.Printf("Nolan's store have [%d] games.\n", len(games))

	// fmt.Printf("%+v\n", games)
	fmt.Printf("%-5s %-10s %-5s    %-15s\n", "ID", "NAME", "PRICE", "GENRE")
	for _, v := range games {
		fmt.Printf("#%-4d %-10q %5d    %-15s\n", v.id, v.name, v.price, v.genre)
	}
}
