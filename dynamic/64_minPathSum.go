package dynamic

import "math"

func minPathSum(grid [][]int) int {
	memo := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		memo[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			memo[i][j] = -1
		}
	}
	return dp_minPathSum(len(grid)-1, len(grid[0])-1, memo, grid)
}

func dp_minPathSum(m, n int, memo [][]int, grid [][]int) int {
	if m < 0 || n < 0 {
		return math.MaxInt32
	}
	if m == 0 && n == 0 {
		return grid[0][0]
	}

	if memo[m][n] != -1 {
		return memo[m][n]
	}
	res := min(dp_minPathSum(m-1, n, memo, grid), dp_minPathSum(m, n-1, memo, grid)) + grid[m][n]
	memo[m][n] = res
	return res
}
