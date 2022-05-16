package _0211115

/**
 * @Author: hxy
 * @Description:
 * @File:  字符串中的第一个唯一字符_test
 * @Date: 2021/11/15 17:44
 */
//@tag [哈希]
func firstUniqChar(s string) int {
	table := make(map[uint8]int,0)
	for i,_ :=range s {
		table[s[i]]++
	}
	for i,_ := range s {

		if x,_ := table[s[i]];x == 1 {
			return i
		}
	}
	return -1
}