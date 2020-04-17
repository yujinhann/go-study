package tree

import (
	"fmt"
	"strings"
)

func Render(end, start int) {
	if end == 0 {
		fmt.Println(genTail(start))
		return
	}

	fmt.Println(genLinePattern(genLinePatternCount(end, true), genLinePatternCount(start, false)))

	end--
	start++

	Render(end, start)
}

func genLinePattern(emptyCount, starCount int) string {
	var sb strings.Builder

	genCountedLine(&sb, emptyCount, " ")
	genCountedLine(&sb, starCount, "T")

	return sb.String()
}

func genCountedLine(sb *strings.Builder, count int, pattern string) strings.Builder {
	for i := 0; i < count; i++ {
		sb.WriteString(pattern)
	}
	return *sb
}

func genLinePatternCount(count int, isEmpty bool) int {
	if isEmpty {
		return count - 1
	} else {
		return count*2 + 1
	}
}

func genTail(length int) string {
	var sb strings.Builder

	for i := 1; i < length; i++ {
		sb.WriteString(" ")
	}

	sb.WriteString("|")

	return sb.String()
}
