package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	b := strings.Repeat("!", len(msg))
	s := b + strings.ToUpper(msg) + b

	fmt.Println(s)
}
