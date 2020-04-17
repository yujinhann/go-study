package main

import (
	"fmt"
	"os"
	"strconv"
)

var pattern string = "|"

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("초기값을 입력해 주세요.")
		return
	}

	for index, num := range args {
		if index == 0 {
			continue
		}
		limit, err := strconv.Atoi(num)

		if err != nil {
			continue
		}

		Paint(limit)
	}
}

func Paint(limit int) {
	if limit < 0 || limit > 20 {
		return
	}

	for i := 0; i < limit; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(pattern)
		}
		fmt.Println()
	}

	for i := 0; i < limit; i++ {
		for j := limit; j > i; j-- {
			fmt.Print(pattern)
		}
		fmt.Println()
	}
}
