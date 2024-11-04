package dynamic
//https://labuladong.online/algo/dynamic-programming/two-views-of-dp/
// 从s的角度穷举
// dp定义一下：dp(s,t,i,j) 等于从s[i:]找t[j:]的次数
// base case1 j到t的末尾，匹配成功
// base case2 s[i:]  比t[j:]短，肯定不匹配
func numDistinct(s string, t string) int {
	memo := MakeMemo(len(s), len(t))
	var dp func(s, t string, i, j int) int
	dp = func(s, t string, i, j int) int {
		// base case1
		if j == len(t) {
			return 1
		}
		if len(s[i:])<len(t[j:]) {
			return 0 
		}
		if memo[i][j]!= -1{
			return memo[i][j]
		}
		res :=0
		if s[i] == t[j]{
			res += dp(s,t,i+1,j) + dp(s,t,i+1,j+1)
		}else{
			res += dp(s,t,i+1,j)
		}
		memo[i][j] = res
		return res
	}
	return dp(s, t, 0, 0)
}

func MakeMemo(row, col int) [][]int {
	memo := make([][]int, row)
	for i, _ := range memo {
		memo[i] = make([]int, col)
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}
	return memo
}
