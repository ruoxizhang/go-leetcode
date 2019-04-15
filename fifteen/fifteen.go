package main

import "fmt"

func main() {
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	result := make([][]int, 0)

	if numsLen == 0 {
		return result
	}

	sortedNums := mergeSort(nums)

	if sortedNums[0] > 0 {
		return result
	}

	if sortedNums[numsLen-1] < 0 {
		return result
	}

	for i := 0; i < numsLen-2; i++ {
		start := i + 1
		end := numsLen - 1
		for start < end {
			tempSum := sortedNums[i] + sortedNums[start] + sortedNums[end]
			if tempSum == 0 {
				result = append(result, []int{sortedNums[i], sortedNums[start], sortedNums[end]})
				start++
				for start < end && sortedNums[start] == sortedNums[start-1] {
					start++
				}
			} else if tempSum > 0 {
				end--
				for start < end && sortedNums[end] == sortedNums[end+1] {
					end--
				}
			} else {
				start++
				for start < end && sortedNums[start] == sortedNums[start-1] {
					start++
				}
			}
		}
		for i+1 < numsLen-2 && sortedNums[i] == sortedNums[i+1] {
			i++
		}
	}
	return result
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
