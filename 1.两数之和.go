/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	exist := make(map[int]int, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		if index, ok := exist[target-nums[i]]; ok {
			return []int{index, i}
		} else {
			exist[nums[i]] = i
		}
	}
	return nil
}

// @lc code=end

