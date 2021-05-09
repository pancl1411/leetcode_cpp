/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start

// 比较麻烦，两个数组的二分查找，total - mid
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total_len := len(nums1) + len(nums2)
	// 奇数个
	if (total_len % 2) == 1 {
		return float64(getKth(nums1, nums2, total_len/2+1))
	}
	// 偶数个，需要取均值
	left := getKth(nums1, nums2, total_len/2)
	right := getKth(nums1, nums2, total_len/2+1)
	return (float64(left) + float64(right)) / 2.0
}

// 注意边界情况，想清楚数值的意义
func getKth(nums1 []int, nums2 []int, k int) int {
	// 分别从数组中取到的元素，做index时需要减一
	select1, select2, half := 0, 0, 0
	for {
		// 一个数组已取完
		if select1 == len(nums1) {
			return nums2[select2+k-1]
		} else if select2 == len(nums2) {
			return nums1[select1+k-1]
		}
		// 剩下的k走一半
		if k > 1 {
			half = k / 2
		} else {
			half = 1
		}
		// 需要避免越界
		index1 := min(select1+half-1, len(nums1)-1)
		index2 := min(select2+half-1, len(nums2)-1)
		if nums1[index1] < nums2[index2] {
			k = k - (index1 + 1 - select1)
			select1 = index1 + 1
			if k == 0 {
				return nums1[index1]
			}
		} else {
			k = k - (index2 + 1 - select2)
			select2 = index2 + 1
			if k == 0 {
				return nums2[index2]
			}
		}
	}
}

func min(f, s int) int {
	if f < s {
		return f
	}
	return s
}

// @lc code=end

