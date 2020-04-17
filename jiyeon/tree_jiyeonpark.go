package main

import "fmt"

func main() {
	// *
	// ***
	// ******
	// ********

	maxCount := 5
	maxStar := maxCount*2 + 1

	for i := 0; i < maxCount; i++ {
		starCount := i*2 + 1
		emptySpace := (maxStar - starCount) / 2

		for j := 0; j < emptySpace; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < starCount; j++ {
			fmt.Print("*")
		}

		for j := 0; j < emptySpace; j++ {
			fmt.Print(" ")
		}

		fmt.Println()
	}
}
