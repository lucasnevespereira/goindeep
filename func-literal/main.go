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

	// Literal func can be assigned to a variable
	biggestNum := func(a, b int) bool {
		return nums[a] > nums[b]
	}
	sort.Slice(nums, biggestNum)
	fmt.Println(nums) // [55 35 15 5]

	// Note: This function uses closure
	// Closure means that the func() inside of biggestNum acts on the variable nums that it's not inside of it's scope.

}
