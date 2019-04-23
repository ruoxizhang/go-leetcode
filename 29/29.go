package main

import "fmt"

func main() {
	fmt.Println()
}

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}
	positive := false
	if dividend > 0 && divisor > 0 {
		positive = true
	} else if dividend < 0 && divisor < 0 {
		positive = true
		dividend = 0 - dividend
		divisor = 0 - divisor
	} else if dividend < 0 && divisor > 0 {
		positive = false
		dividend = 0 - dividend
	} else {
		positive = false
		divisor = 0 - divisor
	}

	sumArr := []int{}
	dividendTemp := dividend
	divisorTemp := divisor
	for dividendTemp >= divisorTemp {
		sumArr = append(sumArr, divisorTemp)
		divisorTemp = divisorTemp + divisorTemp
	}
	dividendTemp = dividend
	result := 0
	for i := len(sumArr) - 1; i >= 0; i-- {
		result = result + result
		if dividendTemp >= sumArr[i] {
			result = result + 1
			dividendTemp = dividendTemp - sumArr[i]
		}
	}
	if !positive {
		result = 0 - result
	}

	if result > 2147483647 {
		return 2147483647
	}

	if result < -2147483648 {
		return -2147483648
	}
	return result
}
