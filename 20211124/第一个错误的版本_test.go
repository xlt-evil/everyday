package _0211124

/**
 * @Author: hxy
 * @Description:
 * @File:  第一个错误的版本_test
 * @Date: 2021/11/24 15:44
 */
//@tag [二分查找]
func firstBadVersion(n int) int {
	low,high := 0,n
	for low <= high {
		mid := low + ((high - low)>>1)
		if !isBadVersion(mid) {
			low = mid + 1
		}else if mid == 0 || !isBadVersion(mid - 1) {
			return mid
		}else {
			high = mid - 1
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	return true
}