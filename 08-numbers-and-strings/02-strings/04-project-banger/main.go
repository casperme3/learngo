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
	"os"
	"strings"
)

// NOTE: You should always pass it at least one argument

func main() {
	msg := os.Args[1]
	l := len(msg)

	s := strings.ToUpper(msg) + strings.Repeat("!", l)

	fmt.Println(s)
}
