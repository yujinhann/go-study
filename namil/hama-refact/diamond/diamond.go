package diamond

import (
	"fmt"
	"strings"
)

func Render(max, index int, aes bool) {
	if index == 0 {
		return
	}

	fmt.Println(getPattern(index))

	if index == max {
		aes = false
	}

	if aes {
		index++
	} else {
		index--
	}

	Render(max, index, aes)
}

func getPattern(len int) string {
	var sb strings.Builder

	for i := 0; i < len; i++ {
		sb.WriteString("ㅓ")
	}

	return sb.String()
}
