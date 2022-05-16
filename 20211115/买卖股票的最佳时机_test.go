package _0211115

import (
	"fmt"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  买卖股票的最佳时机_test
 * @Date: 2021/11/15 14:31
 */

//@tag [动态规划]
func maxProfit(prices []int) int {
	maxPrice,dp := 0,prices[0]
	for i := 1;i < len(prices);i++ {
		// 先判断是否比买入的小
		if dp > prices[i] {
			dp = prices[i]
			continue
		}
		if prices[i] - dp > maxPrice {
			maxPrice = prices[i] - dp
		}
	}
	return maxPrice
}

func Test_maxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}