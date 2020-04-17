package main

func main() {
	max := 9
	for i := 1; i < max; i++ {
		for j := 1; j < i; j++ {
			print("*")
		}
		println()
	}
	for i := max; i > 0; i-- {
		for j := 1; j < i; j++ {
			print("*")
		}
		println()
	}
	println()
}
