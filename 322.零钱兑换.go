/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
// DFS 会超时，用dp
func coinChange(coins []int, amount int) int {
	//
	if amount == 0 {
		return 0
	}
	//
	dp := make([]int, amount+1)
	//
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] == i {
				dp[i] = 1
			} else if coins[j] < i && dp[i-coins[j]] > 0 {
				if dp[i] == 0 {
					dp[i] = dp[i-coins[j]] + 1
				} else {
					dp[i] = min(dp[i], dp[i-coins[j]]+1)
				}
			}
		}
	}
	fmt.Println(dp)
	//
	if dp[len(dp)-1] == 0 {
		return -1
	}
	return dp[len(dp)-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

