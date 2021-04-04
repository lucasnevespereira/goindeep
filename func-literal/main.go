package main

import "fmt"

func main() {
	// Simple definition
	func() {
		fmt.Println("Hello there")
	}()

	// With Args
	func(msg string) {
		fmt.Println(msg)
	}("Hello w/ args")

}
