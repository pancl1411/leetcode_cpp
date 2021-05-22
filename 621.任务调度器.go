/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
//思路参考：https://leetcode-cn.com/problems/task-scheduler/solution/tong-zi-by-popopop/
func leastInterval(tasks []byte, n int) int {
	// special case
	if len(tasks) == 0 {
		return 0
	}
	// 字母计数 + 排序
	counter := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		idx := int(tasks[i] - 'A')
		counter[idx]++
	}
	sort.Slice(counter, func(i, j int) bool {
		return counter[i] > counter[j] // 注意是逆序
	})
	// 可能需要等待的执行时间
	var maxLenCount int = 0
	for i := 0; i < len(counter); i++ {
		if counter[i] == counter[0] {
			maxLenCount++
		} else {
			break
		}
	}
	withSleepTime := (n+1)*(counter[0]-1) + maxLenCount
	// 排满无等待的执行时间
	noSleepTime := len(tasks)
	//
	return max(withSleepTime, noSleepTime)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end