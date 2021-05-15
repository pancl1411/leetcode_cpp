/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	stack  []int
	minVal int
}

//
// 思路：在栈的基础上,记录最小值,若最小值更新,先将old最小值入栈
//
/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		minVal: math.MaxInt32,
	}
}

func (this *MinStack) Push(val int) {
	if val <= this.minVal {
		this.stack = append(this.stack, this.minVal)
		this.minVal = val
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	top := this.Top()
	if top == this.minVal {
		this.stack = this.stack[:len(this.stack)-1]
		this.minVal = this.Top()
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

