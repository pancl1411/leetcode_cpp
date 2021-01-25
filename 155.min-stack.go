/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	data []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{}
	stack.data = make([]int, 0)
	stack.min = math.MaxInt32
	return stack
}

func (this *MinStack) Push(x int) {
	// Tricks: push the old minimum value first,
	// 	while the minimum value changes after pushing the new value x
	if x <= this.min {
		this.data = append(this.data, this.min)
		this.min = x
	}
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	del := this.data[len(this.data)-1]
	// if pop operation could result in the changing of the current minimum value,
	// pop twice and change the current minimum value to the last minimum value.
	if del == this.min {
		this.data = this.data[:len(this.data)-1]
		this.min = this.data[len(this.data)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

