package main

func main() {
	println("")
	star(8)
	println("\n---------\n")
	tree(4)

}

func tree(step int) {
	star := 1
	for n := step; n > 0; n-- {
		for i := 0; i < n; i++ {
			print(" ")
		}
		for j := 1; j <= star; j++ {
			print("*")
		}
		println("")
		star += 2
	}
}

func star(step int) {
	star := 0
	max := 0
	if step%2 == 0 {
		max = step/2 - 1
	} else {
		max = step / 2
	}
	for n := 0; n < step; n++ {
		if n > max {
			star--
		} else {
			star++
		}
		for i := 0; i < star; i++ {
			print("*")
		}
		println("")
	}
}
