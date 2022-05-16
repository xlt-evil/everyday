package _0220322

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  分割等和子集
 * @Date: 2022/03/22 15:54
 */

func Division(num []int) bool {
	var (
		dp []int
		m = map[int]struct{}{}
	)
	s := sum(num)
	if s % 2 != 0 {
		return false
	}
	s = s / 2
	for _,v := range num {
		var tmp []int
		if v == s {
			return true
		}
		for _,x := range dp {
			t := v + x
			if t == s {
				return true
			}else if t > s {
				continue
			}else if _,ok := m[t];!ok {
				tmp = append(tmp,t)
				m[t] = struct{}{}
			}
		}
		if _,ok := m[v];!ok {
			dp = append(dp, v)
		}
		dp = append(dp, tmp...)
	}
	return false
}

func sum(nums []int) int {
	a := 0
	for _,v := range nums {
		a += v
	}
	return a
}

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i] // 表示当前元素 => 可以理解为占用容量需要多少空间
		for j := 1; j <= target; j++ { // j 表示当前背包的容量
			if j >= v { // 说明当前背包可以抓下当前元素
				// j - v 表示 因为可以安装当前元素，所以还要判定，剩余的空间是否能够填满当前包的剩余空间
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}



func Test_Division(t *testing.T) {
	fmt.Println(canPartition([]int{1,5,5,1}))
}



