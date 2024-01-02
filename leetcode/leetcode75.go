package leetcode

/*
75. 颜色分类
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。
链接：https://leetcode.cn/problems/sort-colors/?envType=study-plan-v2&envId=top-100-liked
*/
func sortColors(nums []int) {
	n0, n1 := 0, 0
	for _, v := range nums {
		if v == 0 {
			n0++
		} else if v == 1 {
			n1++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < n0 {
			nums[i] = 0
		} else if i >= n0 && i < n0+n1 {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
