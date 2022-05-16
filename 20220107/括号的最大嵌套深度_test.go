package _0220107

/**
 * @Author: hxy
 * @Description:
 * @File:  括号的最大嵌套深度_test
 * @Date: 2022/01/07 10:44
 */

// 参数一定是有效的括号
func maxDepth(s string) (ans int) {
	size := 0
	for _, ch := range s {
		if ch == '(' {
			size++
			if size > ans {
				ans = size
			}
		} else if ch == ')' {
			size--
		}
	}
	return
}
