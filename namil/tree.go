package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	print("\033[H\033[2J")

	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("초기값을 입력해 주세요.")
		return
	}

	seed, err := strconv.Atoi(arg[1])

	if err != nil {
		fmt.Println("초기 값은 숫자만 입력 가능합니다.")
		return
	}

	if seed < 3 || seed > 40 {
		fmt.Println("초기값이 3~40 이내의 숫자여야 합니다.")
		return
	}

	tree(seed)
}

func tree(seed int) {
	for i := 0; i < seed; i++ {
		for j := 0; j < seed*2-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print(getPattern(j))
		}
		time.Sleep(200 * time.Millisecond)
		fmt.Println()
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < seed*2; j++ {
			fmt.Print(" ")
		}
		time.Sleep(200 * time.Millisecond)
		fmt.Println("/")
	}

	for i := 0; i < seed*4; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Print("/")
	}

	time.Sleep(20 * time.Second)
}

func getPattern(num int) string {
	if num%2 == 0 {
		return "|"
	}
	return "-"
}
