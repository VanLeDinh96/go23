package cmd

import "fmt"

func Execute() {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0}, 
		{0, 0, 0, 0, 0, 0, 0}, 
		{1, 0, 0, 1, 1, 1, 0}, 
		{0, 1, 0, 1, 1, 1, 0}, 
		{0, 1, 0, 0, 0, 0, 0}, 
		{0, 1, 0, 1, 1, 0, 0}, 
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	count := CountRectangles(arr)
	fmt.Printf("%v", count)
}

func CountRectangles(rectangles [][]int) int {
	rows := len(rectangles)
	cols := len(rectangles[0])
	rectangleCount := 0

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || rectangles[row][col] == 0 {
			return
		}

		rectangles[row][col] = 0

		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rectangles[i][j] == 1 {
				dfs(i, j)
				rectangleCount++
			}
		}
	}

	return rectangleCount
}