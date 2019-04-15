package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {

	sAllMap := make(map[int]map[string]int)

	sMapLevel1 := map[string]int{
		"()": 1,
	}
	sAllMap[1] = sMapLevel1

	for i := 2; i <= n; i++ {
		sMapPreLevel := sAllMap[i-1]
		sMapCurLevel := map[string]int{}
		for key := range sMapPreLevel {
			newKey0 := key + "()"
			newKey1 := "()" + key
			newKey2 := "(" + key + ")"
			sMapCurLevel[newKey0] = 1
			sMapCurLevel[newKey1] = 1
			sMapCurLevel[newKey2] = 1
		}
		sAllMap[i] = sMapCurLevel
	}

	result := []string{}
	for key := range sAllMap[n] {
		result = append(result, key)
	}
	return result
}
