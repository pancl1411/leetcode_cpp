/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return quickSearchKth(nums, 0, len(nums)-1, len(nums)-k+1)
}

func quickSearchKth(nums []int, start, end, k int) int {
	// 直接使用less的idx，k保持不更新
	// 在全部相同的情况，通过less-1、less+1逐个逼近
	less := partition(nums, start, end)
	if less+1 == k {
		return nums[less]
	} else if less+1 < k {
		return quickSearchKth(nums, less+1, end, k)
	} else {
		return quickSearchKth(nums, start, less-1, k)
	}
}

// [start,end]
func partition(nums []int, start, end int) int {
	//
	mid := start + (end-start)/2
	key, left, right := nums[mid], start, end
	nums[mid], nums[left] = nums[left], nums[mid]
	//
	for left <= right {
		// move left
		for left <= end && nums[left] <= key {
			left++
		}
		// move right
		for right >= start && nums[right] > key {
			right--
		}
		//
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	//
	nums[start], nums[left-1] = nums[left-1], nums[start]
	//
	return left - 1
}

// @lc code=end

