package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	println("- make() -")
	s := make([]int, 5, 10)
	println(len(s), cap(s))
	fmt.Println(s)

	// s[6] = 5
	//fmt.Println(s)

	println("- Nil Slice -")

	var ss []int
	if ss == nil {
		println("nil!")
	}
	println(len(ss), cap(ss))

	println("- sub slice -")
	b := []int{1, 2, 3, 4, 5, 6}
	b = b[2:5]
	fmt.Println(b)

	println("- append() -")
	c := []int{0, 1}
	fmt.Println(c)
	c = append(c, 2)
	fmt.Println(c)
	c = append(c, 3, 4, 5)
	fmt.Println(c)

	println("- append(), make()")
	d := make([]int, 1, 2)
	fmt.Println(d)
	println(len(d), cap(d))
	d = append(d, 3, 4)
	fmt.Println(d)
	println(len(d), cap(d))

	println("- copy() -")
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	fmt.Println(target)
	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target))

}
