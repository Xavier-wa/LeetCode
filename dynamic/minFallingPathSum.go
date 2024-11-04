package dynamic

import "math"

func minFallingPathSum(matrix [][]int) int {
	rowLen, colLen := len(matrix), len(matrix[0])
	memo := make([][]int, rowLen)
	for i := 0; i < rowLen; i++ {
		memo[i] = make([]int, colLen)
		for j := 0; j < colLen; j++ {
			memo[i][j] = 66666
		}
	}
	//有可能结果打在最后一行的任何一列上
	res := math.MaxInt32
	for j := 0; j < colLen; j++ {
		res = min(res, dp(matrix, rowLen-1, j, memo))
	}
	return res
}

// matrix[0][...]到matrix[n-1][...]的最小路径和
func dp(matrix [][]int, i, j int, memo [][]int) int {
	//索引判断
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return 99999
	}
	//base case
	if memo[i][j] != 66666 {
		return memo[i][j]
	}
	if i == 0 {
		memo[i][j] = matrix[0][j]
		return matrix[0][j]
	}
	//状态转移
	memo[i][j] = matrix[i][j] + min(dp(matrix,i-1,j-1,memo),dp(matrix,i-1,j,memo),dp(matrix,i-1,j+1,memo))
	return memo[i][j]
}
