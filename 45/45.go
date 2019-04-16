package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	if numsLen == 1 {
		return 0
	}

	cacheSteps := make([]int, numsLen)
	furthest := 0
	for i := 0; i < numsLen; i++ {
		if i+nums[i] >= numsLen-1 {
			return cacheSteps[i] + 1
		}

		if i+nums[i] > furthest {
			for j := furthest + 1; j <= i+nums[i]; j++ {
				cacheSteps[j] = cacheSteps[i] + 1
			}
			furthest = i + nums[i]
		}
	}
	return 0
}
