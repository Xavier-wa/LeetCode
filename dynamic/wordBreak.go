package dynamic

func wordBreak(s string, wordDict []string) bool {

	memo := make([]int, len(s))
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

		if memo[i] != 0 {
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
