package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		width  = 50
		height = 10
		buflen = (width*2 + 1) * height //you could also use this len to allocate buf

		cellEmpty = ' '
		cellBall  = '⚾'

		fps       = 16
		speed     = time.Second / fps
		maxFrames = 30 * fps //secs * fps

	)

	var (
		px, py int    //ball position
		vx, vy = 1, 1 //velocities

		cell rune
	)

	board := make([][]bool, height)
	for row := range board {
		board[row] = make([]bool, width)
	}

	buf := make([]rune, 0, height*width*2+height) //or (width*2+1) * height
	// fmt.Printf("00) buf[%p], len(%d), cap(%d)\n", buf, len(buf), cap(buf))

	// Draw a smiley
	// board[2][12] = true
	// board[2][16] = true
	// board[4][14] = true
	// board[5][14] = true
	// board[5][9] = true
	// board[5][19] = true
	// board[6][10] = true
	// board[6][18] = true
	// board[7][12] = true
	// board[7][14] = true
	// board[7][16] = true

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
