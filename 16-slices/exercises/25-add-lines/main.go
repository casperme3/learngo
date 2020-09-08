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
	"strings"

	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Add a newline after each sentence
//
//  You have a slice that contains Beatles' awesome song:
//  Yesterday. You want to add newlines after each sentence.
//
//  So, create a new slice and copy every words into it. Lastly,
//  after each sentence, add a newline character ('\n').
//
//
// ORIGINAL SLICE:
//
//   [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
// EXPECTED SLICE (NEW):
//
//   [yesterday all my troubles seemed so far \n away now it looks as though they are here to stay \n oh i believe in yesterday \n]
//
//
// CURRENT OUTPUT
//
//  yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday
//
// EXPECTED OUTPUT
//
//  yesterday all my troubles seemed so far away
//  now it looks as though they are here to stay
//  oh i believe in yesterday
//
//
// RESTRICTIONS
//
//  + Don't use `append()`, use `copy()` instead.
//
//  + Don't cheat like this:
//
//    fmt.Println(lyric[:8])
//    fmt.Println(lyric[8:18])
//    fmt.Println(lyric[18:23])
//
//  + Create a new slice that contains the sentences
//    with line endings.
//
//
// NOTE
//
// If the program does not work, please update your
// local copy of the prettyslice package:
//
//   go get -u github.com/inancgumus/prettyslice
//
//
// ---------------------------------------------------------

func main() {
	// You need to add a newline after each sentence in another slice.
	// Don't touch the following code.
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday that's just for yesterday yeahh ohhh haaahhh! yeah wadda pack is this song`)

	// 0										  7		8
	// yesterday all my troubles seemed so far away
	// 9										 18  	19
	// now it looks as though they are here to stay
	// 20				   24		 25
	// oh i believe in yesterday
	// 26									32			33
	// that's just for yesterday yeahh ohhh haaahhh!
	// 34				    	 39  40
	// yeah wadda pack is this song
	// ===================================
	//
	// ~~~ CHANGE THIS CODE ~~~
	//
	fix := make([]string, len(lyric)+5)

	cutpoint := [...]int{8, 18, 23, 30, 36}
	for i, n := 0, 0; n < len(lyric); i++ {
		n += copy(fix[n+i:], lyric[n:cutpoint[i]])

		fix[n+i] = "\n"
		fmt.Printf("%d\n", n)
		// if i > 1 {
		// 	break
		// }
	}

	//TODO: use a loop instead dammit!!!!
	// copy(fix, lyric[:8])
	// fix[8] = "\n"
	// copy(fix[9:19], lyric[8:18])
	// fix[19] = "\n"
	// copy(fix[20:25], lyric[18:23])
	// fix[25] = "\n"
	//
	// ===================================

	// Currently, it prints every sentence on the same line.
	// Don't touch the following code.
	s.Show("fix slice", fix)

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}

func init() {
	//
	// YOU DON'T NEED TO TOUCH THIS
	//
	// This initializes some options for the prettyslice package.
	// You can change the options if you want.
	//
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	//
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 5       // prints max 15 elements per line
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}
