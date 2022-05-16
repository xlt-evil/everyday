package _0211209

import "container/list"

/**
 * @Author: hxy
 * @Description:
 * @File:  包含min函数的栈_test
 * @Date: 2021/12/09 16:06
 */

//@tag [设计]
type MinStack struct {
	stack *list.List
	minStack2 *list.List
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: list.New(),
		minStack2: list.New(),
	}
}


func (this *MinStack) Push(x int)  {
	this.stack.PushBack(x)
	if this.minStack2.Len() == 0 || this.minStack2.Back().Value.(int) >= x {
		this.minStack2.PushBack(x)
	}
}


func (this *MinStack) Pop()  {
	if this.stack.Len() == 0 {
		return
	}
	if this.stack.Remove(this.stack.Back()).(int) == this.Min() {
		this.minStack2.Remove(this.minStack2.Back())
	}
	return
}


func (this *MinStack) Top() int {
	e := this.stack.Back()
	return e.Value.(int)
}


func (this *MinStack) Min() int {
	return this.minStack2.Back().Value.(int)
}
