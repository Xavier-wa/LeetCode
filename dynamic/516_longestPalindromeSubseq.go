package dynamic

func longestPalindromeSubseq(s string) int {
	//用 dp[i][j] 表示字符串 s 的下标范围 [i,j] 内的最长回文子序列的长度。假设字符串 s 的长度为 n，则只有当 0≤i≤j<n 时，才会有 dp[i][j]>0，否则 dp[i][j]=0。

	//s[i] == s[j] 就要从中间找最长的子然后加上2
	//s[i] != s[j] 就找i+1:j 和i:j-1 哪个子序列最长
	//搞清楚迭代的方向
	//想要的最终结果是dp[0][len(s)-1]
	//i不可能是从0往len(s)-1迭代的，所以可以找到i是从尾往前迭代的
	//因为要先从短的开始求,所以j不会是从len(s)-1开始迭代,所以从i+1开始迭代
	//这样就可以从base触发，从底到顶迭代
	dp := make([][]int, len(s))
	for i, _ := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1 //base case 单个字符一定是
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}
