package calc

func Sum(i ...int) int {
	sum := 0
	for _, j := range i {
		sum += j
	}
	return sum
}
