package _0211202

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  相对名次_test
 * @Date: 2021/12/02 11:41
 */

//@tag [排序]
func findRelativeRanks(score []int) []string {
	ans := make([]string,len(score))
	hash := map[int]int{}
	for i,_ :=range score {
		hash[score[i]] = i
	}
	sort.Ints(score)
	for i,_ := range score {
		ans[hash[score[i]]] = mapper(len(score) - i)
	}
	return ans
}

func mapper(i int) string {
	switch i {
	case 1: return "Gold Medal"
	case 2: return "Silver Medal"
	case 3: return "Bronze Medal"
	default:
		return strconv.Itoa(i)
	}
}

func Test_findRelativeRanks(t *testing.T) {
	fmt.Println(findRelativeRanks([]int{5,4,3,2,1}))
}