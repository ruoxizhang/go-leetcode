package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	tempResult := []int{}
	find(&candidates, target, &tempResult, 0, &result)
	return result
}

func find(candidates *[]int, sum int, tempResult *[]int, cur int, result *[][]int) {
	if sum == 0 {
		add := make([]int, len(*tempResult))
		copy(add, *tempResult)
		*result = append(*result, add)
	} else {
		for i := cur; i < len(*candidates); i++ {
			value := (*candidates)[i]
			if value > sum {
				return
			}
			*tempResult = append(*tempResult, value)
			find(candidates, sum-value, tempResult, i, result)
			*tempResult = (*tempResult)[:len(*tempResult)-1]
		}
	}
}
