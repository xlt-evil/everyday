package _0211207

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  全排列_test
 * @Date: 2021/12/07 17:04
 */
//@tag [回溯法]
func permute(nums []int) [][]int {
	var (
		temp []int
		ans [][]int
		dfs func([]int)
	)
	dfs = func(tt []int) {
		if len(temp) == len(nums) {
			p := make([]int,len(nums))
			copy(p,temp)
			ans = append(ans,p)
			return
		}
		for i := 0 ;i < len(tt);i++ {
			p := make([]int,len(tt))
			copy(p,tt)
			temp = append(temp,tt[i])
			t := append(p[:i],p[i+1:]...)
			dfs(t)
			temp = temp[:len(temp) - 1]
		}
	}
	dfs(nums)
	return ans
}

func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1,2,3,4,5,6,7,8,9,10,11}))

}