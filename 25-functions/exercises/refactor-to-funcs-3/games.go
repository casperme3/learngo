// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"fmt"
)

const data = `
[
        {
                "id": 1,
                "name": "war god",
                "genre": "adventure",
                "price": 5
        },
        {
                "id": 2,
                "name": "comx 5",
                "genre": "strategy",
                "price": 11
        },
        {
                "id": 3,
                "name": "craftedmind",
                "genre": "sandb",
                "price": 23
        }
]`

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

func load() (games []game, err error) {
	var impt []jsonGame
	if err = json.Unmarshal([]byte(data), &impt); err != nil {
		fmt.Println("Error in  Unmarshal")
		return games, err
	}

	//Store imported data to games slice
	for _, g := range impt {
		games = append(games, game{item{g.ID, g.Name, g.Price}, g.Genre})
	}

	return games, nil
}

func addGame(games []game, g game) []game {
	return append(games, g)
}

func newGame(id, price int, name, genre string) game {
	return game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
}

func indexByID(games []game) (byID map[int]game) {
	byID = make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}
	return
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
