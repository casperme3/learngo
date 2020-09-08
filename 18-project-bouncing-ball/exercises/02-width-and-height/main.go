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

// ---------------------------------------------------------
// EXERCISE: Adjust the width and height automatically
//
//  In this exercise, your goal is simple:
//
//    Instead of setting the width and height manually,
//    you need to get the width and height of the terminal
//    screen from your operating system.
//
//    Don't worry, it is easier than it sounds. You just
//    need to read a few documentation and install a
//    Go package.
//
//  1. Go here: https://godoc.org/golang.org/x/crypto/ssh/terminal
//
//     Download the package:
//     go get -u golang.org/x/crypto/ssh/terminal
//
//  2. Find the function that gives you the width and height
//     of the terminal.
//
//  3. Call that function from your program and get the
//     width and height.
//
//  4. When an error occurs while retrieving the width
//     and height, report it.
//
//  5. Set the width and height of the board.
//
//  6. After solving the above steps, update your program
//     to use my screen package instead. It offers an
//     easier way to get the width and height of a
//     terminal.
//
//     go get -u https://github.com/inancgumus/screen
//
//
// BONUS
//
//  1. When you set the width, you may see that the ball
//     goes beyond the left and right borders. This happens
//     because the ball emoji spans to multiple console
//     columns (or cells). Ordinary characters have a
//     single column.
//
//     1. Get the width of the ball emoji using a function
//        from the following package:
//
//        go get -u github.com/mattn/go-runewidth
//
//     2. Divide the width using the rune width of the
//        ball emoji.
//
//  2. Your terminal may have borders, so reduce the
//     height by taking into account the height of
//     your terminal borders.
//
//
// EXPECTED OUTPUT
//
//  When you run the program, the ball should start
//  animating across the total width and height of your
//  terminal screen dynamically.
//
//  Currently you set width and height manually, so it
//  wasn't matter whether your terminal was bigger or
//  smaller, but now it will be!
//
//
// HINT
//
//  Please take a look at this if you get stuck.
//
//  You need to pass the Standard Out file handler
//  to the function that returns you the dimensions.
//
//  Check out my screen package to find out how I'm
//  passing the Standard Out file handler.
//
//     https://github.com/inancgumus/screen/blob/master/dimensions.go
//
// ---------------------------------------------------------

func main() {
	const (
		// width  = 50
		// height = 10
		// buflen = (width*2 + 1) * height //you could also use this len to allocate buf

		cellEmpty = ' '
		cellBall  = '⚾'

		fps       = 20
		speed     = time.Second / fps
		maxFrames = 30 * fps //secs * fps

	)

	var (
		px, py int    //ball position
		vx, vy = 1, 1 //velocities

		cell rune
	)

	// width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	// if err != nil {
	// 	fmt.Println("ERROR: ", err)
	// 	return
	// }

	width, height := screen.Size()
	// fmt.Printf(">>> %d:%d\n", width, height)

	board := make([][]bool, height)
	for row := range board {
		board[row] = make([]bool, width)
	}

	buflen := (width*2 + 1) * height
	buf := make([]rune, 0, buflen) //or (width*2+1) * height
	// fmt.Printf("00) buf[%p], len(%d), cap(%d)\n", buf, len(buf), cap(buf))

	screen.Clear()

	for loop := 0; loop < maxFrames; loop++ {
		// Clear 'buf' to be reused
		buf = buf[:0]
		//Resets the board to empty
		for y := range board {
			for x := range board[0] {
				board[y][x] = false
			}
		}

		px += vx
		py += vy

		//Controls the direction of the ball either (+)right or (-)left
		if px >= width-1 || px <= 0 {
			vx *= -1
		}
		//Controls the direction of the ball either (+)down or (-)up
		if py >= height-1 || py <= 0 {
			vy *= -1
		}

		// fmt.Printf("[%d][%d]\n", px, py)
		// Set the current location of the ball in the board
		board[py][px] = true

		//Drawing of the board and the ball
		for y := range board {
			for x := range board[0] {
				cell = cellEmpty
				if board[y][x] {
					cell = cellBall
				}
				// fmt.Print(string(cell) + " ")
				buf = append(buf, cell, ' ')
			}
			// fmt.Println()
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Print(string(buf))
		// fmt.Printf("%02d) buf[%p], len(%d), cap(%d)\n", loop+1, buf, len(buf), cap(buf))
		time.Sleep(speed)
	}
	// fmt.Printf("ZZ) buf[%p], len(%d), cap(%d)\n", buf, len(buf), cap(buf))
}
