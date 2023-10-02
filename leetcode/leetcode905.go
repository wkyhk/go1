package leetcode

/*
905. 按奇偶排序数组
给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。

返回满足此条件的 任一数组 作为答案。
链接：https://leetcode.cn/problems/sort-array-by-parity/
*/
func sortArrayByParity(nums []int) (ans []int) {
	for _, v := range nums {
		if v%2 == 0 {
			ans = append([]int{v}, ans...)
		} else {
			ans = append(ans, v)
		}
	}
	return

}
