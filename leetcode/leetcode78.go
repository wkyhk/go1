package leetcode

/*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/
func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for _, v := range nums {
		for _, s := range res {
			sub := make([]int, len(s))
			copy(sub, s)
			sub = append(sub, v)
			res = append(res, sub)
		}
	}
	return res
}
