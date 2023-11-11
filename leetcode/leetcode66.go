package leetcode

/*
66. 加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
链接：https://leetcode.cn/problems/plus-one/description/
*/
func plusOne(digits []int) []int {
	var carry int
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			carry = 0
		}
		if carry == 0 {
			break
		}
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
