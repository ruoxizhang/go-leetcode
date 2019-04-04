package main

import (
	"fmt"
)

func main() {
	fmt.Println("there is a problem fixing here!")
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums1Temp []int
	var nums2Temp []int
	if len(nums1) > len(nums2) {
		nums1Temp = nums2
		nums2Temp = nums1
	} else {
		nums1Temp = nums1
		nums2Temp = nums2
	}

	length1 := len(nums1Temp)
	lenght2 := len(nums2Temp)

	minI := 0
	maxI := length1
	halfLen := (len(nums1) + len(nums2) + 1) / 2

	for minI <= maxI {
		i := (minI + maxI) / 2
		j := halfLen - i
		if i < maxI && nums1Temp[i] < nums2Temp[j-1] {
			minI = i + 1
		} else if i > minI && nums1Temp[i-1] > nums2Temp[j] {
			maxI = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2Temp[j-1]
			} else if j == 0 {
				maxLeft = nums1Temp[i-1]
			} else {
				if nums1Temp[i-1] > nums2Temp[j-1] {
					maxLeft = nums1Temp[i-1]
				} else {
					maxLeft = nums2Temp[j-1]
				}
			}
			if (length1+lenght2)%2 == 1 {
				return float64(maxLeft)
			}
			minRight := 0
			if i == length1 {
				minRight = nums2Temp[j]
			} else if j == lenght2 {
				minRight = nums1Temp[i-1]
			} else {
				if nums1Temp[i] > nums2Temp[j] {
					minRight = nums2Temp[j]
				} else {
					minRight = nums1Temp[i]
				}
			}
			return float64((maxLeft + minRight)) / 2.0
		}
	}

	return 0.0
}
