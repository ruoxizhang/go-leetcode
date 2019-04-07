package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	if x > 2147483647 || x < (-2147483648) {
		return 0
	}
	digits := 1
	tempX := x
	for tempX/10 != 0 {
		tempX = tempX / 10
		digits++
	}

	tempX1 := x
	result := 0
	for i := digits; i > 0; i-- {
		result = result + (tempX1%10)*int(math.Pow10(i-1))
		tempX1 = tempX1 / 10
	}

	if result > 2147483647 || result < (-2147483648) {
		return 0
	}
	return result
}
