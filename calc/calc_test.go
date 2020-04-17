package calc

import "testing"

func Test_Sum(t *testing.T) {
	v0 := Sum(1, 2, 3)
	if v0 != 6 {
		t.Fatal("1+2+3 == 6")
	}

	v1 := Sum(6, 5)
	if v1 != 11 {
		t.Fatal("5 + 6 == 11")
	}
}
