/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	answer := make([][]int, 0)
	// special case
	if len(nums) < 3 {
		return answer
	}
	// sort, a<=b<=c
	sort.Ints(nums)
	// 思路参考：https://leetcode-cn.com/problems/3sum/solution/pai-xu-shuang-zhi-zhen-zhu-xing-jie-shi-python3-by/
	for first := 0; first < len(nums); first++ {
		// a >0, a+b+c>0
		if nums[first] > 0 {
			break
		}
		//  a去重
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}
		// 固定a, 求解b+c = -a
		second, third := first+1, len(nums)-1
		for second < third {
			// b去重(a固定，b去重后，c一定不是重复的)
			if second != first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			// 双指针
			if nums[first]+nums[second]+nums[third] < 0 {
				second++
			} else if nums[first]+nums[second]+nums[third] > 0 {
				third--
			} else {
				// ==0 , CPU分支预测，概率小的写下面
				answer = append(answer, []int{nums[first], nums[second], nums[third]})
				second++
			}
		}
	}
	//
	return answer
}

// @lc code=end

