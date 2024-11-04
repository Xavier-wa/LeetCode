package dynamic

// dp[i] 是以num[i]为结尾的最长递增子序列的长度
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := -1
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}

func lengthOfLISJoke(nums []int) int {
    top := make([]int, len(nums))
    // 牌堆数初始化为 0
    var piles int
    for i := 0; i < len(nums); i++ {
        // 要处理的扑克牌
        poker := nums[i]

        // ***** 搜索左侧边界的二分查找 *****
        var left, right int = 0, piles
        for left < right {
            mid := (left + right) / 2
            if top[mid] > poker {
                right = mid
            } else if top[mid] < poker {
                left = mid + 1
            } else {
                right = mid
            }
        }
        // ********************************
        
        // 没找到合适的牌堆，新建一堆
        if left == piles {
            piles++
        }
        // 把这张牌放到牌堆顶
        top[left] = poker
    }
    // 牌堆数就是 LIS 长度
    return piles
}