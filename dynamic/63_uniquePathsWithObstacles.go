package dynamic

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memo := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		memo[i] = make([]int, len(obstacleGrid[0]))
		for j := 0; j < len(obstacleGrid[0]); j++ {
			memo[i][j] = -1
		}
	}
	return dp_uniquePathsWithObstacles(len(obstacleGrid)-1, len(obstacleGrid[0])-1, memo, obstacleGrid)
}

func dp_uniquePathsWithObstacles(m, n int, memo [][]int, obstacleGrid [][]int) int {
	if m < 0 || n < 0 {
		return 0
	}
	if m == 0 && n == 0 {
		if obstacleGrid[0][0] == 1 { //要考虑 [[1]]这种情况
			return 0
		}
		return 1
	}
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	if obstacleGrid[m][n] == 1 {
		memo[m][n] = 0
		return 0
	}
	memo[m][n] = dp_uniquePathsWithObstacles(m-1, n, memo, obstacleGrid) + dp_uniquePathsWithObstacles(m, n-1, memo, obstacleGrid)
	return memo[m][n]
}
