/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	//
	queue := newMaxQueue(nums)
	for i := 0; i < k-1; i++ {
		queue.Push(i)
	}
	//
	answer := make([]int, len(nums)-k+1)
	for i := k - 1; i < len(nums); i++ {
		queue.Push(i)
		queue.PopLess(i - k)
		answer[i-k+1] = queue.GetMax()
	}
	return answer
}

// 实现最大队列
// 思路参考单调队列：https://leetcode-cn.com/problems/sliding-window-maximum/solution/hua-dong-chuang-kou-zui-da-zhi-by-leetco-ki6m/
type maxQueue struct {
	nums        []int // 存的值
	windowIndex []int // 存的窗口索引, queue
}

func newMaxQueue(nums []int) *maxQueue {
	return &maxQueue{
		nums:        nums,
		windowIndex: make([]int, 0),
	}
}

// 入队时,如果前面value较小，可以直接删除其idx
func (q *maxQueue) Push(idxOf int) {
	for len(q.windowIndex) > 0 && q.nums[idxOf] >= q.nums[q.windowIndex[len(q.windowIndex)-1]] {
		q.windowIndex = q.windowIndex[:len(q.windowIndex)-1]
	}
	q.windowIndex = append(q.windowIndex, idxOf)
}

// 窗口左侧合法性判断
func (q *maxQueue) PopLess(idxLess int) {
	for len(q.windowIndex) > 0 && q.windowIndex[0] <= idxLess {
		q.windowIndex = q.windowIndex[1:]
	}
}

func (q *maxQueue) GetMax() int {
	return q.nums[q.windowIndex[0]]
}

// @lc code=end

