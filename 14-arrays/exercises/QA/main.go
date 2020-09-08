package main

import "fmt"

func main() {
	type (
		threeA [3]int
		threeB [3]int
	)

	nums := threeA{1, 2, 3}
	nums2 := threeA(threeB{1, 2, 3})

	fmt.Println(nums == nums2)
}
