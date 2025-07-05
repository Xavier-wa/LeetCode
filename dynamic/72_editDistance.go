package dynamic

// 解决两个字符串的动态规划问题，一般都是用两个指针 i, j 分别指向两个字符串的头部或尾部，然后尝试写状态转移方程。
func editDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i, _ := range memo {
		memo[i] = make([]int, len(word2))
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp_editDistance(word1, word2, len(word1)-1, len(word2)-1, memo)
}
func dp_editDistance(s1, s2 string, i, j int, memo [][]int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s1[i] == s2[j] {
		memo[i][j] = dp_editDistance(s1, s2, i-1, j-1, memo)
	} else {
		memo[i][j] = min(
			dp_editDistance(s1, s2, i, j-1, memo),   //+
			dp_editDistance(s1, s2, i-1, j-1, memo), //r
			dp_editDistance(s1, s2, i-1, j, memo),
		) + 1
	}
	return memo[i][j]
}

func editDistance2(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1) //dp[i][j]代表 word[0...i-1] 变成 word2[0...j-1]的最小次数
	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ { // 其中一个是空串，那么最小次数就是另一个字符串的长度
		dp[i][0] = i //为什么dp长度要比word1多一位,考虑了空串的情况相当于 word1[0...i-1] 变成 word2[0...-1]而不是word1[0...0]
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ { //穷举
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] //相等不需要操作
			} else {
				dp[i][j] = min( //取最小再加上一次操作
					dp[i][j-1],
					dp[i-1][j],
					dp[i-1][j-1],
				) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
