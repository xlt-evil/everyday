package main

/**
 * @Author: hxy
 * @Description:
 * @File:  两数之和
 * @Date: 2022/05/16 11:22
 */


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i,_ := range nums {
		if v,ok := m[target - nums[i]];ok {
			return []int{v,i}
		}else {
			m[nums[i]] = i
		}
	}
	return nil
}