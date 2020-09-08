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

//     2- Add a func that creates and returns a game.
//
//        Name  : newGame
//        Inputs: id, price, name and genre
//        Output: game
func newGame(id, price int, name, genre string) game {
	return game{item{id, name, price}, genre}
}

//     3- Add a func that adds a `game` to `[]game` and returns `[]game`:
//
//        Name  : addGame
//        Inputs: []game, game
//        Output: []game
func addGame(games []game, g game) []game {
	return append(games, g)
}

//     4- Add a func that uses newGame and addGame funcs to load up the game
//        store:
//
//        Name  : load
//        Inputs: none
//        Output: []game
func load() (games []game) {
	gow := newGame(1, 51, "god of wara", "action")
	xxx := newGame(2, 42, "xxx-com3", "strategy")
	mc := newGame(3, 33, "minedurcraft", "sandboxed")

	games = addGame(games, gow)
	games = addGame(games, xxx)
	games = addGame(games, mc)

	// fmt.Printf("load: %v\n", games)
	return
}

//     5- Add a func that indexes games by ID:
//
//        Name  : indexByID
//        Inputs: []game
//        Output: map[int]game
//
func indexByID(games []game) (toID map[int]game) {
	toID = make(map[int]game)
	for _, g := range games {
		toID[g.id] = g
	}
	return
}

//     6- Add a func that prints the given game:
//
//        Name  : printGame
//        Inputs: game
//
func printGame(g game) {
	fmt.Printf("#%d: %-20q %-20s $%d\n", g.id, g.name, "[["+g.genre+"]]", g.price)
}
