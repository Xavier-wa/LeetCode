package dynamic

//dp[k] 代表偷k间房子的最佳结果

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[1]

	for k := 2; k <= len(nums); k++ {
		dp[k] = max(dp[k-1], dp[k-2]+nums[k-1])
	}
	return dp[len(nums)]
}
