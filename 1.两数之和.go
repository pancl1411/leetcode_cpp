/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

// map保存已有值，target-nums[i]查找
func twoSum(nums []int, target int) []int {
	exist := make(map[int]int, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		if index, ok := exist[target-nums[i]]; !ok {
			//CPU分支预测，高概率条件写前面
			exist[nums[i]] = i
		} else {
			return []int{index, i}
		}
	}
	return nil
}

// @lc code=end

