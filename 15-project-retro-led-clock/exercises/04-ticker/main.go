// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Ticker: Slide the Clock
//
//  Your goal is slide the placeholders every second.
//  Please run the solution to see it in action.
//
//
//  THIS IS A HARD EXERCISE:
//  + It will take you days but it will worth it.
//  + For experienced developers, this can take an hour or so.
//
//
//  1. You need to determine the starting and the ending digits to create
//     the sliding effect.
//
//
//  2. Each second, start from the next placeholder, skip the previous one.
//     This means: Only draw the next placeholders.
//
//     Like this:
//
//     12:40:31
//     2:40:31
//     40:31
//     0:31
//     :31
//     31
//     1
//
//
//  3. After the last placeholder is displayed, fill the lines for the missing
//     placeholders, and then start from the first placeholder. Draw it to the
//     right part of the screen.
//
//     Like this:
//
//     12:40:31
//     2:40:31
//     40:31
//     0:31
//     :31
//     31
//     1
//            1
//           12
//          12:
//         12:4
//        12:40
//       12:40:
//      12:40:3
//     12:40:31
//
//     As you can see, you need to draw the clock from the right part of the
//     screen, beginning from the first placeholder.
//
//
// HINTS
//   + You would need to clear the screen inside the loop instead of once.
//     Otherwise the previous placeholders will be left on the screen.
//
//
// EXPECTED OUTPUT
//   Please run the solution to see it in action. Do not look at the source-code
//   though.
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
	maxMove   = 8
	maxSlider = 16
)

func main() {
	// screen.Clear()
	//Infinite Loop: updates every 1 second
	var (
		counter int
	)

	// maintask:
	for {
		screen.Clear()
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]blocktype{
			blank, blank, blank, blank,
			blank, blank, blank, blank,
		}

		tmpClock := [...]blocktype{
			digits[hour/10], digits[hour%10],
			tickOn,
			digits[min/10], digits[min%10],
			tickOn,
			digits[sec/10], digits[sec%10],
		}

		if (sec % 2) == 0 {
			tmpClock[sepa1], tmpClock[sepa2] = blank, blank
		}

		clkLen := len(clock) - 1
		slide := counter % maxSlider
		moveI := slide % maxMove
		// fmt.Printf("count: %d\n", )

		switch {
		case slide < maxMove: //0~7
			for i := range clock {
				if (i + moveI) >= maxMove {
					break
				}
				clock[i] = tmpClock[i+slide]
			}
		case slide > maxMove: //9~16
			for i := 0; i < moveI; i++ {
				clock[clkLen-i] = tmpClock[moveI-1-i]
			}
		}

		//Drawing of clock in the terminal
		for line := range clock[0] {
			for num := range clock {
				fmt.Print(clock[num][line], "  ")
			}
			fmt.Println()
		}

		counter++
		time.Sleep(time.Second)
	}
}
