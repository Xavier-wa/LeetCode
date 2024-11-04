package dynamic

import "math"

// 打家劫舍数组特点是相邻的不能同时选择
// 也就是num[i]拿了就不可以选nums[i+1]
// 目前的情况则是 nums[i] 拿了 nums[i]-1 和nums[i] + 1 就不能拿了

// 排序 可以-1 和 +1 的情况变成一种就是+1
// 但是更好的方法是把原nums 的元素当成新数组的下标志，新数组则是统计该下标在原数组出现的次数
// 这样就把问题转成了 num[i]选了就不可以选nums[i+1]

// 再用打家劫舍的方法设置dp数组
// 定义 dp 数组：dp[i] 表示偷到第 i 个房间最多能偷多少钱
// basecase dp[0] = 0,dp[1] = newNums[1]
// 递归：dp[i] = max(dp[i-1], dp[i-2]+i*rob[i])

// 转换成功后，屋子的金额是下标出现的次数 * 下标
func deleteAndEarn(nums []int) int {
	// 定义转换为打家劫舍的数组，因为这道题只和数字的值有关系

	NumsCount := make([]int, maxInArray(nums)+1)
	for _, v := range nums {
		NumsCount[v]++ //v 是原数组的元素，作为新数组的索引，新数组的元素为索引的count
	}
	if len(NumsCount) == 1 {
		// 原数组nums 没有元素
		return 0
	}
	if len(NumsCount) == 2 {
		return NumsCount[1] // 原数组元素只有一个元素
	}
	dp := make([]int, len(NumsCount))
	dp[0], dp[1] = 0, NumsCount[1]
	for i := 2; i < len(NumsCount); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+i*NumsCount[i])
	}
	return dp[len(dp)-1]
}

func maxInArray(nums []int) int {
	last := math.MinInt32
	for _, v := range nums {
		if v > last {
			last = v
		}
	}
	return last
}
