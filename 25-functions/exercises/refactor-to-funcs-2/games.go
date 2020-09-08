// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func newGame(id, price int, name, genre string) game {
	return game{item{id, name, price}, genre}
}

func addGame(games []game, g game) []game {
	return append(games, g)
}

func load() (games []game) {
	gow := newGame(1, 51, "god of wara", "action")
	xxx := newGame(2, 42, "xxx-com3", "strategy")
	mc := newGame(3, 33, "minedurcraft", "sandboxed")

	games = addGame(games, gow)
	games = addGame(games, xxx)
	games = addGame(games, mc)

	return
}

func indexByID(games []game) (toID map[int]game) {
	toID = make(map[int]game)
	for _, g := range games {
		toID[g.id] = g
	}
	return
}

func printGame(g game) {
	fmt.Printf("#%d: %-20q %-20s $%d\n", g.id, g.name, "[["+g.genre+"]]", g.price)
}
