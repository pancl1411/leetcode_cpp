/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
// 思路参考：https://leetcode-cn.com/problems/queue-reconstruction-by-height/solution/xian-pai-xu-zai-cha-dui-dong-hua-yan-shi-suan-fa-g/
func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 {
		return nil
	}
	// 排序（先身高逆序，后前面人数正序）
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	//fmt.Println(people)
	// 因为是按照身高从高到矮，所以此时idx就是应该插入的位置
	answer := make([][]int, 0)
	for i := 0; i < len(people); i++ {
		idx := people[i][1]
		// slice 中间插入方法
		answer = append(answer[:idx], append([][]int{people[i]}, answer[idx:]...)...)
	}
	return answer
}

// @lc code=end

