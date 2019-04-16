package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	// 1.     1
	// 2.     11
	// 3.     21
	// 4.     1211
	// 5.     111221
	previousResult := []int{1, 1}
	for i := 2; i < n; i++ {
		newResult := []int{}
		counter := 1
		tempByte := previousResult[0]
		for j := 1; j < len(previousResult); j++ {
			if previousResult[j] == tempByte {
				counter++
			} else {
				newResult = append(newResult, int(counter))
				newResult = append(newResult, tempByte)
				tempByte = previousResult[j]
				counter = 1
			}
		}
		newResult = append(newResult, int(counter))
		newResult = append(newResult, tempByte)
		previousResult = newResult
	}

	stringResult := ""
	for _, val := range previousResult {
		stringResult = stringResult + strconv.Itoa(val)
	}
	return stringResult
}
