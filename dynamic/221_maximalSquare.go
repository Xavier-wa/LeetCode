package dynamic

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)
	res := 0

	for i, _ := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	//dp[i-1][j-1]是以matix[i-1][j-1]为右下角的正方形的最大边长
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res * res
}
