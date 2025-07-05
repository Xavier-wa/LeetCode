package dynamic

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	//dp[i][j]表示s[j:i]是否为回文串if
	dp := make([][]int, len(s))
	maxLen := 1
	maxS, maxE := 0, 0
	for i, _ := range dp {
		dp[i] = make([]int, len(s))
	}
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1] == 1) { //i-j==0,i-j==1,i-j==2 分别代表，奇数长度最中间字符，偶数长度最中间两个字符，奇数长度最中间两个字符
				dp[j][i] = 1
				if i-j+1 > maxLen {
					maxLen = i - j + 1
					maxS = j
					maxE = i
				}
			}
		}
	}
	return s[maxS : maxE+1]
}
