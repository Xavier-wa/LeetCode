package dynamic

func uniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	return dp_uniquePaths(m-1, n-1, memo)
}

func dp_uniquePaths(m, n int, memo [][]int) int {
	if m == 0 || n == 0 {
		return 1
	}
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	res := dp_uniquePaths(m-1, n, memo) + dp_uniquePaths(m, n-1, memo)
	memo[m][n] = res
	return res
}
//因为大小为mxn 所以最多只能走到m-1,n-1
// dp(m,n) = 到达dp(m,n)有多少条路径
// dp(m,n-1) + 1,dp(m-1,n) +1
