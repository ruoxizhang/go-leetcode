package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}

	var runeMap map[rune]int
	runeMap = make(map[rune]int)
	start := 0
	end := 0
	longest := 0

	for index, char := range s {
		if indexExist, ok := runeMap[char]; ok {
			if start <= indexExist {
				if longest < index-start {
					longest = index - start
				}
				start = indexExist + 1
			}
		}
		runeMap[char] = index
		end = index
	}

	if end-start+1 > longest {
		return end - start + 1
	}

	return longest
}
