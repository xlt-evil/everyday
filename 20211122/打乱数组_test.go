package _0211122

import "math/rand"

/**
 * @Author: hxy
 * @Description:
 * @File:  打乱数组
 * @Date: 2021/11/22 14:20
 */
//@tag [其他]
//KnuthKnuth 洗牌算法

type Solution struct {
	nums, original []int
}


func Constructor(nums []int) Solution {
	return Solution{ nums,append([]int(nil),nums...)}
}


func (this *Solution) Reset() []int {
	copy(this.nums,this.original)
	return this.nums
}


func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	for i := range this.nums {
		j := i + rand.Intn(n - i)
		this.nums[i],this.nums[j] = this.nums[j],this.nums[i]
	}
	return this.nums
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */