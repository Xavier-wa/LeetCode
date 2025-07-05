package dynamic

func wordBreak(s string, wordDict []string) bool {

	memo := make([]int, len(s))
	for i, _ := range memo {
		memo[i] = -1
	}
	wordDictMap := make(map[string]bool)
	for _, v := range wordDict {
		wordDictMap[v] = true
	}
	//dp 定义 s[i:]是否可以在wordDict中找到
	var dp func(s string, i int) bool
	dp = func(s string, i int) bool {
		if i == len(s) {
			return true
		}

		if memo[i] != -1 {
			return memo[i] == 1
		}
		for l := 1; i+l <= len(s); l++ {
			sub := s[i : i+l]
			if wordDictMap[sub] && dp(s, i+l) {
				memo[i] = 1
				return true
			}
		}
		memo[i] = 0
		return false
	}
	return dp(s, 0)
}

// 迭代
func wordBreakPush(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w)) //记录最长，s[j:i] 不能长过maxLen
	}

	n := len(s)
	f := make([]bool, n+1) //把空字符串作为basecase f[0] 所以是n+1
	f[0] = true
	for i := 1; i <= n; i++ { // 从1开始 f[i]求 到f[n]
		for j := i - 1; j >= max(i-maxLen-1, 0); j-- { //s[j:i] 不能长过maxLen 所以j的下限是max(i-maxLen-1, 0)
			if f[j] && words[s[j:i]] {
				f[i] = true //要求f[i] 即s[:i]是否在wordDict中 就是s[:i] 之间可以拆成f[j]和s[j:i] 并且f[j]是true,s[j:i]在wordDict中
				break
			}
		}
	}
	return f[n]
}
