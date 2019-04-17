package main

import "fmt"

func main() {
	nextPermutation([]int{1, 3, 2})
}

func nextPermutation(nums []int) {
	lenNums := len(nums)
	if lenNums < 2 {
		return
	}
	swapped := false
	end := 0
Exit:
	for i := lenNums - 1; i >= 0; i-- {
		for j := lenNums - 1; j > i; j-- {
			if nums[i] < nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
				end = i + 1
				swapped = true
				break Exit
			}
		}
	}

	if swapped {
		findNSwap(&nums, lenNums-1, end)
	}

	if !swapped {
		for i := 0; i < lenNums/2; i++ {
			temp := nums[i]
			nums[i] = nums[lenNums-1-i]
			nums[lenNums-1-i] = temp
		}
	}

	fmt.Printf("%v", nums)
}

func findNSwap(nums *[]int, startIndex int, endIndex int) bool {
	if endIndex >= startIndex {
		return true
	}
	for i := startIndex; i > endIndex; i-- {
		for j := i - 1; j >= endIndex; j-- {
			if (*nums)[i] < (*nums)[j] {
				temp := (*nums)[i]
				(*nums)[i] = (*nums)[j]
				(*nums)[j] = temp
				findNSwap(nums, startIndex, j+1)
			}
		}
	}
	return true
}
