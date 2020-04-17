package main

import "fmt"

func main() {
	// *
	// **
	// ***
	// **
	// *

	var maxCount int
	fmt.Print("값 입력(3 이상) : ")
	fmt.Scanln(&maxCount)
	if maxCount > 2 {
		printHalfDiamond(maxCount)
	}
}

func printHalfDiamond(maxCount int) {
	middle := maxCount / 2
	for count := 0; count < maxCount; count++ {
		if count < middle {
			for sign := 0; sign <= count; sign++ {
				fmt.Print("*")
			}
		} else {
			if maxCount%2 == 0 && count == middle {
				continue
			}
			for sign := 0; sign < maxCount-count; sign++ {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func jy() {
	maxCount := 9
	for i := 0; i < maxCount; i++ {
		if i < maxCount/2 {
			for k := 0; k <= i; k++ {
				fmt.Print("*")
			}
		} else {
			for k := 0; k < maxCount-i; k++ {
				fmt.Print("*")
			}
		}

		fmt.Println()
	}
}
