package main

import (
	"./diamond"
	"./tree"
	"fmt"
)

func main() {
	print("\033[H\033[2J")
	fmt.Println("=== HAMA님의 코드 리팩토링한 버전입니다===")
	fmt.Print("\n\n1. 크리스마스 트리\n")
	tree.Render(16, 0)

	fmt.Print("\n\n2. 다이아몬드\n")
	diamond.Render(16, 1, true)
}
