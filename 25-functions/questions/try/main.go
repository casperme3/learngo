package main

import "fmt"

// IT SHOULD PRINT: [10 5 2]
func main() {
	stats := make([]int, 0, 2)
	// stats := []int{10, 5}
	fmt.Printf("%d. %d\n", cap(stats), len(stats))
	stats = append(stats, 10, 5)
	add(stats, 2)
	fmt.Print(stats)
}

func add(stats []int, n int) {
	stats = append(stats, n)
	fmt.Printf("add: %d\n", stats)
	fmt.Printf("%d. %d\n", cap(stats), len(stats))
}
