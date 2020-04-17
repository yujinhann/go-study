package main

import "fmt"

func main() {
	mapScore := scoreMap()
	names, score := makeSlice(mapScore)
	printFindResult(getScore(names, score, "yujin"))
	names, score = add(names, score, "namil", 100)
	delete(names, score, "yujin")

}

func scoreMap() map[string]int {
	score := map[string]int{
		"yujin":    80,
		"hama":     95,
		"jiyeon":   90,
		"hyeonsul": 85,
	}
	return score
}

func makeSlice(scoreMap map[string]int) (names []string, score []int) {
	for key, value := range scoreMap {
		names = append(names, key)
		score = append(score, value)
	}
	return
}

func getScore(names []string, score []int, findName string) (name string, findScore int, exists bool) {
	// 인덱스 가져오기
	index := getIndex(names, findName)
	// 인덱스로 val, exists return 하기
	name = findName
	if index == -1 {
		exists = false
	} else {
		findScore = score[index]
		exists = true
	}
	return
}

func getIndex(names []string, name string) (index int) {
	index = -1
	for i, value := range names {
		if value == name {
			index = i
			break
		}
	}
	return
}

func add(names []string, score []int, addName string, addScore int) ([]string, []int) {
	names = append(names, addName)
	score = append(score, addScore)
	return names, score
}

func delete(names []string, score []int, delName string) ([]string, []int) {
	index := getIndex(names, delName)
	if index == -1 {
		return names, score
	} else {
		copy(names, append(names[0:index], names[index+1:]...))
		copy(score, append(score[0:index], score[index+1:]...))
		return names[:len(names)-1], score[:len(score)-1]
	}
}

func printSlice(names []string, score []int) {
	fmt.Println(names)
	fmt.Println(score)
}

func printFindResult(name string, score int, exists bool) {
	if exists {
		fmt.Printf("%s : %d점 \n", name, score)
	} else {
		fmt.Println("해당하는 이름을 가진 학생이 없습니다.")
	}
}
