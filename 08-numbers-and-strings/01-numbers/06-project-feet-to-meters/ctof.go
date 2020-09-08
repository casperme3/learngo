package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	cel, _ := strconv.ParseFloat(os.Args[1], 64)
	far := cel*1.8 + 32

	fmt.Printf("%g celcius is equal to %g fahrenheit.\n", cel, far)

	var n float64 = 1
	n++
	fmt.Println(n)
}
