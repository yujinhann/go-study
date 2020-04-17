package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	println(arr[2])

	println("----")

	var b = [3]int{1, 2, 3}
	for index, item := range b {
		println(index, item)
	}

	println("---")

	var c = [...]int{1, 2}
	for i, item := range c {
		println(i, item)
	}

	println("---")

	var multiArr [2][2]int
	multiArr[0][0] = 1
	println(multiArr[0][0])

	println("---")

	var multiArr2 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	println(multiArr2[0][2])

	println("---")

	s := []int{1, 2, 3}
	a := make([]int, 1, 2)
	copy(a, s)
	fmt.Println(a)
}
