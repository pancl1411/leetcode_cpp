/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 课程入度表
	indegree := make([]int, numCourses)
	// 依赖课程的map, 用于快速找到某个前置课程完成时, 哪些课程的入度表需要更新
	preMap := make(map[int][]int, 0)
	// 入度表、依赖map更新
	for i := 0; i < len(prerequisites); i++ {
		//
		indegree[prerequisites[i][0]]++
		//
		v, ok := preMap[prerequisites[i][1]]
		if ok {
			preMap[prerequisites[i][1]] = append(v, prerequisites[i][0])
		} else {
			preMap[prerequisites[i][1]] = []int{prerequisites[i][0]}
		}
	}
	// 可以完成的课程栈
	finishStack := make([]int, 0)
	finishCount := 0
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			finishStack = append(finishStack, i)
			finishCount++
		}
	}
	//
	for len(finishStack) != 0 {
		// pop
		top := finishStack[len(finishStack)-1]
		finishStack = finishStack[:len(finishStack)-1]
		//
		courses, ok := preMap[top]
		if ok {
			for _, cour := range courses {
				indegree[cour]--
				if indegree[cour] <= 0 {
					finishStack = append(finishStack, cour)
					finishCount++
				}
			}
		}
	}
	//
	if finishCount == numCourses {
		return true
	}
	return false

}

// @lc code=end

