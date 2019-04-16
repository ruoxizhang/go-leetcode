package main

import "fmt"

func main() {
	fmt.Printf("%v", permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1}))
}

func permuteUnique(nums []int) [][]int {
	newNums := mergeSort(nums)
	result := [][]int{}
	backTrace(newNums, []int{}, &result)
	return filterNums(result)
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

func filterNums(result [][]int) [][]int {
	newResult := [][]int{}
	for _, val := range result {
		if !isContains(newResult, val) {
			newResult = append(newResult, val)
		}
	}
	return newResult
}

func isContains(result [][]int, temp []int) bool {
	for _, val := range result {
		if isArrayEquals(val, temp) {
			return true
		}
	}
	return false
}

func isArrayEquals(nums1 []int, nums2 []int) bool {
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func mergeSort(nums []int) []int {
	oldNums := nums
	numsLen := len(oldNums)
	for i := 1; i < numsLen; i = i * 2 {
		newNums := make([]int, numsLen)
		n := 0
		for j := 0; j <= numsLen; j = 2*i + j {
			k := j
			m := j + i
			for k < i+j && m < 2*i+j && k < numsLen && m < numsLen {
				if oldNums[k] < oldNums[m] {
					newNums[n] = oldNums[k]
					n++
					k++
				} else {
					newNums[n] = oldNums[m]
					n++
					m++
				}
			}

			for k < i+j && k < numsLen {
				newNums[n] = oldNums[k]
				n++
				k++
			}

			for m < i*2+j && m < numsLen {
				newNums[n] = oldNums[m]
				n++
				m++
			}
		}
		oldNums = newNums
	}
	return oldNums
}
