package main

import "fmt"

type com struct {
	brand string
}

func main() {
	var a, b *com

	fmt.Print(a == b)

	a = &com{"Apple"}
	b = &com{"Apple"}
	fmt.Print(" ", a == b, " ", *a == *b)
}
