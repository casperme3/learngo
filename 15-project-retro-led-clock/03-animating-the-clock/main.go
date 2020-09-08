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
	"time"

	"github.com/inancgumus/screen"
)

type blocktype [5]string

func main() {
	separator := blocktype{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	digits := [...]blocktype{
		{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		},
		{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		},
	}

	screen.Clear()
	//Infinite Loop: updates every 1 second
	for {
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		// fmt.Println(now)

		clock := [...]blocktype{
			digits[hour/10], digits[hour%10],
			separator,
			digits[min/10], digits[min%10],
			separator,
			digits[sec/10], digits[sec%10],
		}

		screen.MoveTopLeft()
		for line := range clock[0] {
			for num := range clock {
				fmt.Print(clock[num][line], " ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
