package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{2, 3, 10, 5, 7, 8, 9}))
}

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1

	max := 0
	for start+1 <= end {
		if height[start] > height[end] {
			area := height[end] * (end - start)
			if area > max {
				max = area
			}
			end--
		} else {
			area := height[start] * (end - start)
			if area > max {
				max = area
			}
			start++
		}
	}
	return max
}
