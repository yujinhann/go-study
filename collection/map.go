package main

import "fmt"

func main() {
	println("- nil map -")
	var idMap map[int]string
	if idMap == nil {
		println("true")
	}

	println("- map for make() -")
	idMap = make(map[int]string)
	fmt.Println(idMap)
	idMap[900] = "apple"
	idMap[1] = "grape"
	idMap[1] = "banana"
	println(idMap[900], idMap[1], idMap[2])
	delete(idMap, 900)
	println(idMap[900], idMap[1], idMap[2])

	println("- map 초기화 for literal -")
	tickers := map[string]string{
		"one":   "im one",
		"two":   "im two",
		"three": "im three",
	}
	val, exists := tickers["one"]
	println(val, ", ", exists)
	val2, exists2 := tickers["four"]
	println(val2, ", ", exists2)

	m := make(map[int]string)
}
