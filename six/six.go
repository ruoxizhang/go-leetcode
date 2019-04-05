package main

import "fmt"

func main() {
	fmt.Println(convert("abcd", 3))
}

func convert(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}

	stringLen := len(s)
	lenPerGroup := numRows*2 - 2
	groups := stringLen/lenPerGroup + 1
	resultCharArr := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		for j := 0; j < groups; j++ {
			if j*lenPerGroup+i < stringLen {
				resultCharArr = append(resultCharArr, s[j*lenPerGroup+i])
			}

			if i != 0 && i != numRows-1 {
				if j*lenPerGroup+lenPerGroup-i < stringLen {
					resultCharArr = append(resultCharArr, s[j*lenPerGroup+lenPerGroup-i])
				}
			}
		}
	}

	return string(resultCharArr)
}
