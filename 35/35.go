package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v", "")
}

func searchInsert(nums []int, target int) int {

	nL := len(nums)

	if nL == 0 {
		return 0
	}
	if nL == 1 {
		if target > nums[0] {
			return 1
		}
		return 0
	}

	start := 0
	end := nL - 1
	if target < nums[0] {
		return 0

	}

	if target > nums[nL-1] {
		return nL
	}

	for start < end-1 {
		middle := (end-start+1)/2 + start
		if target > nums[middle] {
			start = middle
		} else {
			end = middle
		}
	}

	if target == nums[start] {
		return start
	}

	return end
}
