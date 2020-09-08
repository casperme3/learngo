package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	in.Scan()
	in.Scan()

	fmt.Println(in.Text())
}
