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
)

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {
	scientists := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Nikola", "Tesla", "teslabomb"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
	}

	for i, name := range scientists {
		fmt.Printf("%-15s%-15s%-15s\n", name[0], name[1], name[2])

		if i == 0 {
			fmt.Printf("%s\n", strings.Repeat("=", 40))
		}
	}

	// scientists := [...][3]string{
	// 	{"Nikola", "Tesla", "ACpower"},
	// 	{"Albert", "Einstein", "time"},
	// 	{"Isaac", "Newton", "apple"},
	// 	{"Stephen", "Hawking", "blackhole"},
	// 	{"Marie", "Curie", "radium"},
	// }

	// fmt.Printf("%-15s%-15s%-15s\n", "First Name", "Last Name", "Nickname")
	// fmt.Printf("%s\n", strings.Repeat("=", 40))

	// for _, scientist := range scientists {
	// 	for _, val := range scientist {
	// 		fmt.Printf("%-15s", val)
	// 	}
	// 	fmt.Println()
	// }
}
