package tree

import (
	"strings"
	"testing"
)

func TestGenLinePatternCountForPattern(t *testing.T) {
	v0 := genLinePatternCount(5, false)
	if v0 != 11 {
		t.Fatal("5 * 2 + 1 == 11")
	}
}

func TestGenLinePatternCountForEmpty(t *testing.T) {
	v0 := genLinePatternCount(5, true)
	if v0 != 4 {
		t.Fatal("5 - 1 == 4")
	}
}

func TestGenLine(t *testing.T) {
	var sb strings.Builder
	v0 := genLine(&sb, 5, "-")
	if v0.String() != "-----" {
		t.Fatal("5 : -----")
	}
}
