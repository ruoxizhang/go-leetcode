package main

import "fmt"

func main() {
	fmt.Printf("%v", searchRange([]int{5, 7, 7, 8, 8, 10, 11}, 8))
}

func searchRange(nums []int, target int) []int {
	lenNums := len(nums)

	if lenNums == 0 {
		return []int{-1, -1}
	}

	if lenNums == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	if target > nums[lenNums-1] {
		return []int{-1, -1}
	}
	if target < nums[0] {
		return []int{-1, -1}
	}

	resultStart := findTheFirst(nums, target)
	if resultStart != -1 {
		resultEnd := findTheLast(nums, target)
		return []int{resultStart, resultEnd}
	} else {
		return []int{-1, -1}
	}

	return []int{-1, -1}
}

func findTheFirst(nums []int, target int) int {
	lenNums := len(nums)

	start := 0
	end := lenNums - 1

	for start < end-1 {
		if nums[start] == target && (start == 0 || nums[start-1] != target) {
			return start
		}

		middle := (end-start+1)/2 + start
		if target > nums[middle] {
			start = middle
		} else {
			end = middle
		}
	}

	if nums[start] == target && (start == 0 || nums[start-1] != target) {
		return start
	}
	if nums[end] == target && (end == 0 || nums[end-1] != target) {
		return end
	}

	return -1
}

func findTheLast(nums []int, target int) int {
	lenNums := len(nums)

	start := 0
	end := lenNums - 1

	for start < end-1 {
		if nums[end] == target && (end == lenNums-1 || nums[end+1] != target) {
			return end
		}

		middle := (end-start+1)/2 + start
		if target < nums[middle] {
			end = middle
		} else {
			start = middle
		}
	}

	if nums[end] == target && (end == lenNums-1 || nums[end+1] != target) {
		return end
	}
	if nums[start] == target && (start == lenNums-1 || nums[start+1] != target) {
		return start
	}

	return -1
}
