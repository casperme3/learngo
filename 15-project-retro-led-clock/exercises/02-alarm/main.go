// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Set an Alarm
//
//  Goal is printing " ALARM! " every 10 seconds.
//
// EXPECTED OUTPUT
//
//     ██   ███       ███  ██        ███  ███
//      █     █   ░     █   █    ░     █  █ █
//      █   ███       ███   █        ███  ███
//      █   █     ░     █   █    ░     █    █
//     ███  ███       ███  ███       ███  ███
//
//          ███  █    ███  ██   █ █   █
//          █ █  █    █ █  █ █  ███   █
//          ███  █    ███  ██   █ █   █
//          █ █  █    █ █  █ █  █ █
//          █ █  ███  █ █  █ █  █ █   █
//
//     ██   ███       ███  ██        █ █  ██
//      █     █   ░     █   █    ░   █ █   █
//      █   ███       ███   █        ███   █
//      █   █     ░     █   █    ░     █   █
//     ███  ███       ███  ███         █  ███
//
// HINTS
//
//  <<< BEWARE THE SPOILER! >>>
//
//  I recommend you to first solve the exercise yourself before reading the
//  following hint.
//
//
//  + You can create a new array named alarm with the same length of the
//    clocks array, so you can copy the alarm array to the clocks array
//    every 10 seconds.
//
// ---------------------------------------------------------

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const (
	sepa1 = iota + 2
	_
	_
	sepa2
)

func main() {
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

		if sec%10 == 0 {
			clock = alarm
		} else if (sec % 2) == 0 {
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
