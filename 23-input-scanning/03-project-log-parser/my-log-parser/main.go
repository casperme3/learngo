package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		in      = bufio.NewScanner(os.Stdin)
		stats   = make(map[string]int)
		domains []string
		line    int
	)
	for in.Scan() {
		line++
		field := strings.Fields(in.Text())
		if len(field) < 2 {
			fmt.Printf("Missing input at line:%d\n", line)
			return
		}

		domain := field[0]
		visits, err := strconv.Atoi(field[1])
		if err != nil || visits < 0 {
			fmt.Printf("Wrong number value at line: %d\n", line)
			return
		}

		// fmt.Printf("%s || %d\n", domain, visits)

		if _, ok := stats[domain]; ok {
			stats[domain] += visits
			continue
		}
		stats[domain] = visits
		domains = append(domains, domain)
	}

	fmt.Printf("%-30s %8s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)
	for _, d := range domains {
		fmt.Printf("%-30s %8d\n", d, stats[d])
	}

	fmt.Printf("map len[%d][%d]\n", len(stats), len(domains))

}
