package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly #2
//
//   This challenge is based on the previous Moodly challenge.
//
//   1. Display the usage if the username or the mood is missing
//
//   2. Change the moods array to a multi-dimensional array.
//
//      So, create two inner arrays:
//        1. One for positive moods
//        2. Another one for negative moods
//
//   4. Randomly select and print one of the mood messages depending
//      on the given mood command-line argument.
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name] [positive|negative]
//
//   go run main.go Socrates
//     [your name] [positive|negative]
//
//   go run main.go Socrates positive
//     Socrates feels good 👍
//
//   go run main.go Socrates positive
//     Socrates feels happy 😀
//
//   go run main.go Socrates positive
//     Socrates feels awesome 😎
//
//   go run main.go Socrates negative
//     Socrates feels bad 👎
//
//   go run main.go Socrates negative
//     Socrates feels sad 😞
//
//   go run main.go Socrates negative
//     Socrates feels terrible 😩
// ---------------------------------------------------------
const (
	usage = "[your name] [positive|negative]"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	name, mood := args[0], args[1]
	mods := [...][3]string{
		{"happy 😀", "good 👍", "great 😎"},
		{"sad", "bad", "terrible"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(mods[0]))
	fmt.Printf(">%d\n", n)
	if mood == "positive" {
		mood = mods[0][n]
	} else if mood == "negative" {
		mood = mods[1][n]
	} else {
		fmt.Printf("Invalid mood!\n")
		return
	}

	fmt.Printf("%s feels %s.\n", name, mood)
}
