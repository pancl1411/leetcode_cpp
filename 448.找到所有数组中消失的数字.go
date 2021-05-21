/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	// 思想：数字对应的idx位置加一个大数,第二轮遍历时看对应位置小于大数的就是
	for _, v := range nums {
		idx := (v - 1) % n
		nums[idx] += n
	}
	//
	result := make([]int, 0)
	for i, v := range nums {
		if v <= n {
			result = append(result, i+1)
		}
	}
	return result
}

// @lc code=end

