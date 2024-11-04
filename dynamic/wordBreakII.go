package dynamic

func wordBreakII(s string, wordDict []string) []string {

	memo := make([][]string, len(s))
	wordDictMap := make(map[string]bool)
	for _, v := range wordDict {
		wordDictMap[v] = true
	}
	//dp 定义 s[i:]是否可以在wordDict中找到
	var dp func(s string, i int) []string
	dp = func(s string, i int) []string {
		res := []string{} //这次的结果
		if i == len(s) {
			res = append(res, "")
			return res
		}

		if len(memo[i]) > 0 {
			return memo[i]
		}

		for l := 1; i+l <= len(s); l++ {
			prefix := s[i : i+l]
			if wordDictMap[prefix] {
				nextRes := dp(s, i+l)
				for v := range nextRes {
					if nextRes[v] == "" {
						res = append(res, prefix)
					} else {
						res = append(res, prefix+" "+nextRes[v])
					}
				}
			}
		}
		memo[i] = res
		return res
	}
	return dp(s, 0)
}
