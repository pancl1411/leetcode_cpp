/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
var answer [][]int

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	//
	answer = [][]int{}
	backtrace(nums, []int{}, 0)
	return answer
}

func backtrace(nums, selected []int, start int) {
	//
	if len(nums) == start {
		cp := make([]int, len(selected))
		copy(cp, selected)
		answer = append(answer, cp)
	}
	// tricks : 利用swap和start idx隔离已选/未选的值（或者加个map记录）
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrace(nums, append(selected, nums[start]), start+1)
		nums[i], nums[start] = nums[start], nums[i]
	}
}

// @lc code=end

