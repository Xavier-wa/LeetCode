package dynamic

func tribonacci(n int) int {
	memo := make(map[int]int)
	return dp(n, memo)
}

func dp(n int, memo map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	res := dp(n-1, memo) + dp(n-2, memo) + dp(n-3, memo)
	memo[n] = res
	return memo[n]
}
