/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
// 思路参考：https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/0-1-bei-bao-wen-ti-xiang-jie-zhen-dui-ben-ti-de-yo/
//
func canPartition(nums []int) bool {
	// special case
	if len(nums) <= 1 {
		return false
	}
	//
	var sum int = 0
	for _, v := range nums {
		sum += v
	}
	// special case : sum % 2 ==1
	if (sum & 1) == 1 {
		return false
	}
	half := sum / 2
	//
	dp := make([][]bool, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, half+1)
	}
	//
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 {
				dp[0][j] = (j == nums[0])
			} else if dp[i-1][j] || nums[i] == j {
				dp[i][j] = true
			} else if nums[i] <= j {
				dp[i][j] = dp[i-1][j-nums[i]]
			}
		}
	}
	//
	return dp[len(dp)-1][len(dp[0])-1]
}

// @lc code=end

