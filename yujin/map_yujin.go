package main

import (
	"fmt"
)

func main() {
	// 1. map[string]int 형식의 성적 map을 만듦
	// ex) 우리팀["hama"] = 80
	// 2. 데이터 초기화
	// 3. slice로 index 를 넣었을 때 value를 가져오기
	// ex) string array, int array
	// 4. slice add, remove
	fmt.Println("--- 점수 등록하기 ---")
	count := setStudentCount()
	scoreMap := make(map[string]int)
	enrollScore(count, scoreMap)
	names, scores := makeSlice(scoreMap)
	fmt.Println("--- 점수 찾기 ---")
	fmt.Println("점수 :", findScore(scores, findIndex(names, inputFindName())))
	fmt.Println("--- 학생 추가 ---")
	names, scores = addScore(names, scores)
	fmt.Println("--- 점수 찾기 ---")
	fmt.Println("점수 :", findScore(scores, findIndex(names, inputFindName())))
}

func setStudentCount() (count int) {
	fmt.Print("학생 수 : ")
	fmt.Scanln(&count)
	return
}

func inputNameAndScore() (name string, score int) {
	fmt.Print("이름 : ")
	fmt.Scanln(&name)
	fmt.Print("점수 : ")
	fmt.Scanln(&score)
	return
}

func enrollScore(count int, scoreMap map[string]int) {
	if count == 0 {
		fmt.Println("등록 완료!")
		return
	}
	name, score := inputNameAndScore()
	fmt.Println("----")
	scoreMap[name] = score
	count--
	enrollScore(count, scoreMap)
}

func makeSlice(scoreMap map[string]int) (nameSlice []string, scoreSlice []int) {
	for key, value := range scoreMap {
		nameSlice = append(nameSlice, key)
		scoreSlice = append(scoreSlice, value)
	}
	return
}

func inputFindName() (name string) {
	fmt.Print("이름 : ")
	fmt.Scanln(&name)
	return
}

func findIndex(names []string, findName string) (findIndex int) {
	for index, value := range names {
		if value == findName {
			findIndex = index
			break
		}
		if index == len(names) && value != findName {
			findIndex = -1
		}
	}
	return
}

func findScore(scores []int, findIndex int) int {
	if findIndex == -1 {
		return 0
	} else {
		return scores[findIndex]
	}
}

func addScore(nameSlice []string, scoreSlice []int) ([]string, []int) {
	name, score := inputNameAndScore()
	nameSlice = append(nameSlice, name)
	scoreSlice = append(scoreSlice, score)
	fmt.Println("추가 완료!")
	return nameSlice, scoreSlice
}

func removeScore()
