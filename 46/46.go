package main

import "fmt"

func main() {
	fmt.Printf("%v", permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	result := [][]int{}
	backTrace(nums, []int{}, &result)
	return result
}

func backTrace(nums []int, temp []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, temp)
	} else {
		for i, val := range nums {
			newTemp := append(temp, val)
			newNums := make([]int, len(nums))
			copy(newNums, nums)
			newNums = append(newNums[:i], newNums[i+1:]...)
			backTrace(newNums, newTemp, result)
		}
	}
}
