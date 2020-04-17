package main

import "fmt"

func main() {
	// *
	// **
	// ***
	// **
	// *

	// 맘에 안드는 방법..
	maxCount := 9
	for i := 0; i < maxCount; i++ {
		if i < maxCount/2 {
			for k := 0; k <= i; k++ {
				// 흠..
				fmt.Print("*")
			}
		} else {
			for k := 0; k < maxCount-i; k++ {
				// 흠..
				fmt.Print("*")
			}
		}

		fmt.Println()
	}
}
