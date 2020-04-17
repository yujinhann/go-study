package main

func main() {
	starTest(10)
}

func starTest(size int) {
	i := 1
	top := false
	for i > 0 {
		star := ""
		for j := 0; j < i; j++ {
			star += "*"
		}
		top = top || i == size
		if !top {
			i++
		} else {
			i--
		}
		println(star)
	}
}
func treeTest(size int) {
	cellSize := size*2 - 1
	rootline := ""
	for i := 1; i <= size; i++ {
		treeline := ""
		starSize := 2*i - 1
		emptySize := (cellSize - starSize) / 2
		for empty := 0; empty < emptySize; empty++ {
			treeline += " "
			if i == 1 {
				rootline += " "
			}
		}
		for star := 0; star < starSize; star++ {
			treeline += "*"
			if i == 1 {
				rootline += "|"
			}
		}
		println(treeline)
	}
	println(rootline)
}
