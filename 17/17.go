package main

import "fmt"

func main() {
	fmt.Printf("%v", letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	letterMap := make(map[byte][]byte, 8)
	letterMap['2'] = []byte{'a', 'b', 'c'}
	letterMap['3'] = []byte{'d', 'e', 'f'}
	letterMap['4'] = []byte{'g', 'h', 'i'}
	letterMap['5'] = []byte{'j', 'k', 'l'}
	letterMap['6'] = []byte{'m', 'n', 'o'}
	letterMap['7'] = []byte{'p', 'q', 'r', 's'}
	letterMap['8'] = []byte{'t', 'u', 'v'}
	letterMap['9'] = []byte{'w', 'x', 'y', 'z'}
	result := []string{}
	tempResult := []byte{}
	composeLetters([]byte(digits), &result, &tempResult, &letterMap)
	return result
}

func composeLetters(digits []byte, result *[]string, tempResult *[]byte, letterMap *map[byte][]byte) {
	if len(digits) == 0 {
		*result = append(*result, string(*tempResult))
	} else {
		digit := (digits)[0]
		for _, value := range (*letterMap)[digit] {
			*tempResult = append(*tempResult, value)
			composeLetters((digits)[1:], result, tempResult, letterMap)
			*tempResult = (*tempResult)[:len(*tempResult)-1]
		}
	}
}
