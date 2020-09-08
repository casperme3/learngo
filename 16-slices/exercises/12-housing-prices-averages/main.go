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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		loc                                   []string
		size, beds, baths, prices             []int
		sumSize, sumBeds, sumBaths, sumPrices int
	)

	rows := strings.Split(data, "\n")
	for _, row := range rows {
		col := strings.Split(row, separator)
		for i, v := range col {
			num, _ := strconv.Atoi(v)
			switch i {
			case 0:
				loc = append(loc, v)
			case 1:
				size = append(size, num)
			case 2:
				beds = append(beds, num)
			case 3:
				baths = append(baths, num)
			case 4:
				prices = append(prices, num)
			}
		}
	}

	hdr := strings.Split(header, separator)
	for _, h := range hdr {
		fmt.Printf("%-15s", h)
	}
	fmt.Println("\n" + strings.Repeat("=", 70))
	for i := range rows {
		fmt.Printf("%-15s %-14d %-15d %-12d %-14d\n",
			loc[i], size[i], beds[i], baths[i], prices[i],
		)
		sumSize += size[i]
		sumBeds += beds[i]
		sumBaths += baths[i]
		sumPrices += prices[i]
	}

	denum := float64(len(rows))
	// denum := len(rows)
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Printf("%-15s %-14.2f %-15.2f %-12.2f %-14.2f\n",
		"", float64(sumSize)/denum, float64(sumBeds)/denum, float64(sumBaths)/denum, float64(sumPrices)/denum,
	)
	// fmt.Printf("%-15s %-14.2f %-15.2f %-12.2f %-14.2f\n",
	// 	"", float64(sumSize/denum), float64(sumBeds/denum), float64(sumBaths/denum), float64(sumPrices/denum),
	// )
}
