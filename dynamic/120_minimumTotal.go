package dynamic

// dp[m][n] 为 从下到上到达dp[m][n] 最短路径的长度
// 地推公式 dp[m][n] = min(dp[m+1][n], dp[m+1][n+1]) + triangle[m][n] 每个节点向下有两个节点
// 自底向上迭代
func minimumTotal(triangle [][]int) int {
	h := len(triangle)

	dp := make([][]int, len(triangle))
	for i, _ := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	//从底向上的动态规划
	//base case = triangle[h-1][j]
	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == h-1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}
