package main

import (
	"fmt"
	"sort"
)

func main() {
	// Simple definition
	func() {
		fmt.Println("Hello there")
	}()

	// With Args
	func(msg string) {
		fmt.Println(msg)
	}("Hello w/ args")

	// Use case to sort a list
	nums := []int{35, 55, 5, 15}
	sort.Slice(nums, func(a,b int) bool {
		return nums[a] < nums[b]
	})
	fmt.Println(nums) // [5 15 35 55]


}
