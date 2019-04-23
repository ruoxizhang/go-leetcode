package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	find(&candidates, target, 0, &[]int{}, &result)
	return result
}

func find(candidates *[]int, target int, currentPosi int, tempResult *[]int, result *[][]int) {
	if target == 0 {
		add := make([]int, len(*tempResult))
		copy(add, *tempResult)
		*result = append(*result, add)
	} else {
		used := make(map[int]int)

		for i := currentPosi; i < len(*candidates); i++ {
			value := (*candidates)[i]
			if value > target {
				return
			}

			if _, ok := (used)[value]; !ok {
				(used)[value]++
			} else {
				continue
			}

			*tempResult = append(*tempResult, value)
			find(candidates, target-value, i+1, tempResult, result)
			*tempResult = (*tempResult)[:len(*tempResult)-1]
		}
	}
}
