package leetcode

/*
1550. 存在连续三个奇数的数组
给你一个整数数组 arr，请你判断数组中是否存在连续三个元素都是奇数的情况：如果存在，请返回 true ；否则，返回 false 。
链接：https://leetcode.cn/problems/three-consecutive-odds/
*/
func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1]%2 == 1 && arr[i]%2 == 1 && arr[i+1]%2 == 1 {
			return true
		}
	}
	return false
}
