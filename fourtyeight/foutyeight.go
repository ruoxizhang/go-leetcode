package main

import "fmt"

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix) == 1 {
		return
	}

	if len(matrix[0]) == 0 {
		return
	}
	var currentX int
	var currentY int
	var onTheMove int
	var nextX int
	var nextY int
	var temp int
	var k int
	var j int

	n := len(matrix)
	for i := 0; i <= n/2; i++ {
		for j = i; j < n-i-1; j++ {
			currentX = i
			currentY = j
			onTheMove = matrix[i][j]
			for k = 0; k < 4; k++ {
				nextY = n - 1 - currentX
				nextX = currentY
				temp = matrix[nextX][nextY]
				matrix[nextX][nextY] = onTheMove
				onTheMove = temp
				currentX = nextX
				currentY = nextY
			}
		}
	}
	fmt.Printf("%v", matrix)
}
