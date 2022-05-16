package _0211116

/**
 * @Author: hxy
 * @Description:
 * @File:  赎金信_test
 * @Date: 2021/11/16 11:21
 */

//@tag [哈希]
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[uint8]int,0)
	for i,_ := range magazine {
		m[magazine[i]]++
	}
	for i,_ := range ransomNote {
		if count,_ := m[ransomNote[i]];count > 0 {
			m[ransomNote[i]]--
		}else {
			return false
		}
	}
	return true
}