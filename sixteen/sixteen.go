package main

import "fmt"

func main() {
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 3))
}

func threeSumClosest(nums []int, target int) int {
	numsLen := len(nums)

	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}

	if numsLen == 2 {
		return nums[0] + nums[1]
	}
	if numsLen == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sortedNums := mergeSort(nums)
	var difference = nums[0] + nums[1] + nums[2] - target

	for i := 0; i < numsLen-2; i++ {
		start := i + 1
		end := numsLen - 1
		for start < end {
			tempDifference := sortedNums[i] + sortedNums[start] + sortedNums[end] - target
			if tempDifference == 0 {
				return target
			}
			if abs(tempDifference) < abs(difference) {
				difference = tempDifference
			}

			if tempDifference > 0 {
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
	return target + difference
}

func abs(num int) int {
	if num < 0 {
		return 0 - num
	}
	return num
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
