package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi("-9223372036854775808"))
}

func myAtoi(str string) int {
	numberRuneMap := make(map[rune]int)
	numberRuneMap['1'] = 1
	numberRuneMap['2'] = 2
	numberRuneMap['3'] = 3
	numberRuneMap['4'] = 4
	numberRuneMap['5'] = 5
	numberRuneMap['6'] = 6
	numberRuneMap['7'] = 7
	numberRuneMap['8'] = 8
	numberRuneMap['9'] = 9
	numberRuneMap['0'] = 0
	maxInt := 2147483647
	minInt := -2147483648

	result := 0

	minus := false
	start := false

	for _, value := range str {
		if start == false {
			if !(value == ' ') {
				if value == '-' {
					minus = true
					start = true
				} else if value == '+' {
					start = true
				} else if number, ok := numberRuneMap[value]; ok {
					result = number
					start = true
				} else {
					return 0
				}
			}
		} else {
			if number, ok := numberRuneMap[value]; ok {
				if minus == false && result >= maxInt {
					return maxInt
				} else if minus == true && (0-result) <= minInt {
					return minInt
				}
				result = result*10 + number
			} else {
				if minus == true {
					if (0 - result) >= minInt {
						return 0 - result
					}
					return minInt
				}
				if result <= maxInt {
					return result
				}
				return maxInt
			}
		}
	}

	if minus == true {
		if (0 - result) >= minInt {
			return 0 - result
		}
		return minInt
	}

	if result <= maxInt {
		return result
	}
	return maxInt
}
