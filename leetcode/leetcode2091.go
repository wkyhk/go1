package leetcode

import "fmt"

/*
2091.从数组中移除最大值和最小值
给你一个下标从 0 开始的数组 nums ，数组由若干 互不相同 的整数组成。
nums 中有一个值最小的元素和一个值最大的元素。分别称为 最小值 和 最大值 。你的目标是从数组中移除这两个元素。
一次 删除 操作定义为从数组的 前面 移除一个元素或从数组的 后面 移除一个元素。
返回将数组中最小值和最大值 都 移除需要的最小删除次数。
链接：https://leetcode.cn/problems/removing-minimum-and-maximum-from-array
*/
func minimumDeletions(nums []int) int {
	i, j := 0, 0
	for p, num := range nums {
		if num < nums[i] {
			i = p
		}
		if num > nums[j] {
			j = p
		}
	}
	if i > j {
		i, j = j, i // 保证 i <= j
	}
	return min(min(j+1, len(nums)-i), i+1+len(nums)-j)
}

func TestLeetCode2091() {
	fmt.Println(minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6}))
	fmt.Println(minimumDeletions([]int{0, -4, 19, 1, 8, -2, -3, 5}))
	fmt.Println(minimumDeletions([]int{101}))
}
