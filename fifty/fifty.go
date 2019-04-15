package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(myPow(2.1, 3))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n == -1 {
		return 1 / x
	}

	positive := true
	if n < 0 {
		positive = false
		n = 0 - n
	}

	binaryStr := strconv.FormatInt(int64(n), 2)
	binaryLen := len(binaryStr)

	tempResult := 1.0
	tempIdemResult := x
	for i := binaryLen - 1; i > 0; i-- {
		tempIdemResult = tempIdemResult * tempIdemResult

		if binaryStr[i-1] == '1' {
			tempResult = tempIdemResult * tempResult
		}
	}

	if binaryStr[binaryLen-1] == '1' {
		tempResult = tempResult * x
	}

	if !positive {
		return 1 / tempResult
	}

	return tempResult
}
