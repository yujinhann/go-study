package main

import "fmt"

func main() {
	line := 5
	fmt.Println("- 삼각형 -")
	for i := 0; i < line; i++ {
		for j := line - 1; j > i; j-- {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
