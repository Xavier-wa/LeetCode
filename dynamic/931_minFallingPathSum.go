package dynamic

// dp[i][j]等于经过i,j点的路径的最小和，这是一个子问题，然后从上到下遍历一下，再在最后一行取最小结果
func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i, _ := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == len(matrix[0])-1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + matrix[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			}
		}
	}
	minN := dp[len(matrix)-1][0]
	for j := 1; j < len(matrix[0]); j++ {
		if minN > dp[len(matrix)-1][j] {
			minN = dp[len(matrix)-1][j]
		}
	}
	return minN
}
