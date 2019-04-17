package main

import "fmt"

func main() {
	fmt.Printf("%v", permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1}))
}

func permuteUnique(nums []int) (result [][]int) {
	find(nums, 0, &result)
	return
}

func find(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		*result = append(*result, append([]int{}, nums...))
	}
	freq := make(map[int]bool, len(nums))
	for i := index; i < len(nums); i++ {
		if _, ok := freq[nums[i]]; ok {
			continue
		}
		nums[index], nums[i] = nums[i], nums[index]
		find(nums, index+1, result)
		nums[index], nums[i] = nums[i], nums[index]
		freq[nums[i]] = true
	}
}

func permuteUniqueOld(nums []int) [][]int {
	newNums := mergeSort(nums)
	result := [][]int{}
	backTraceOld(newNums, []int{}, &result)
	return result
}

func backTraceOld(nums []int, temp []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, temp)
	} else {
		usedNums := make(map[int]int, len(nums))
		for i, val := range nums {
			if _, ok := usedNums[val]; !ok {
				usedNums[val] = 1
				newTemp := append(temp, val)
				newNums := make([]int, len(nums))
				copy(newNums, nums)
				newNums = append(newNums[:i], newNums[i+1:]...)
				backTraceOld(newNums, newTemp, result)
			}
		}
	}
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
