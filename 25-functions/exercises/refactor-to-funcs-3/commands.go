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
	"strconv"
	"strings"
)

func runCmd(input string, games []game, byID map[int]game) bool {
	fmt.Println()

	b := true
	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return b
	}

	switch cmd[0] {
	case "quit":
		b = cmdQuit()
	case "list":
		b = cmdList(games)
	case "id":
		b = cmdByID(cmd, games, byID)
	case "save":
		b = cmdSave(games)
	}
	return b
}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

func cmdByID(cmd []string, games []game, byID map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := byID[id]
	if !ok {
		fmt.Println("sorry. i don't have the game")
		return true
	}

	printGame(g)

	return true
}

//        Name  : cmdSave
//        Inputs: []game
//        Output: bool
func cmdSave(games []game) (b bool) {
	var saved []jsonGame
	for _, v := range games {
		saved = append(saved, jsonGame{v.id, v.name, v.genre, v.price})
	}

	out, err := json.MarshalIndent(saved, "", "\t")
	if err != nil {
		fmt.Println(err)
		return b
	}

	fmt.Println(string(out))
	return b
}
