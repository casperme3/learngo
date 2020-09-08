package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Must provide a query word.")
		return
	}
	query := args[0]

	words := make(map[string]bool)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		// fmt.Println(in.Text())
		word := strings.ToLower(in.Text())
		if words[word] || len(word) < 3 {
			continue
		}
		words[word] = true
	}

	fmt.Printf("Map-Size: %d\n", len(words))

	if words[query] {
		fmt.Printf(`The word '%s' is found!`, query)
	} else {
		fmt.Printf(`The word '%s' is not in shakespare poem!`, query)
	}

}
