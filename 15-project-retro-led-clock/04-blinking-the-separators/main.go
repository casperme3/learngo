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

const (
	sepa1 = iota + 2
	_
	_
	sepa2
)

func main() {
	tickOn := blocktype{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	tickOff := blocktype{
		"   ",
		"   ",
		"   ",
		"   ",
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
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		// fmt.Println(now)

		clock := [...]blocktype{
			digits[hour/10], digits[hour%10],
			tickOn,
			digits[min/10], digits[min%10],
			tickOn,
			digits[sec/10], digits[sec%10],
		}

		if (sec % 2) == 0 {
			clock[sepa1], clock[sepa2] = tickOff, tickOff
		}

		for line := range clock[0] {
			for num := range clock {
				fmt.Print(clock[num][line], "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
