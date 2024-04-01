package leetcode

/*
2529. 正整数和负整数的最大计数
给你一个按 非递减顺序 排列的数组 nums ，返回正整数数目和负整数数目中的最大值。

换句话讲，如果 nums 中正整数的数目是 pos ，而负整数的数目是 neg ，返回 pos 和 neg二者中的最大值。
注意：0 既不是正整数也不是负整数。
*/
func maximumCount(nums []int) int {
	pos, neg := 0, 0
	for _, v := range nums {
		if v > 0 {
			pos++
		} else if v < 0 {
			neg++
		}
	}
	if pos > neg {
		return pos
	}
	return neg
}
