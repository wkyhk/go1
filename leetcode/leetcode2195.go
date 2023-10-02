package leetcode

import (
	"fmt"
	"sort"
)

/*
2195. 向数组中追加 K 个整数
给你一个整数数组 nums 和一个整数 k 。请你向 nums 中追加 k 个 未 出现在 nums 中的、互不相同 的 正 整数，并使结果数组的元素和 最小 。

返回追加到 nums 中的 k 个整数之和。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/append-k-integers-with-minimal-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minimalKSum(nums []int, k int) int64 {
	var res int64 = 0
	res = int64((1 + k) * k / 2)
	sort.Ints(nums)
	last := k
	if nums[0] > k {
		return res
	} else {
		last++
		res = res - int64(nums[0]-last)
	}
	for i := 1; i < len(nums) && nums[i] <= last; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		last++
		res = res - int64(nums[i]-last)
	}
	return res

}
func Example2195() {
	arr := []int{1, 4, 25, 10, 25}
	fmt.Println(minimalKSum(arr, 2))
	arr1 := []int{5, 6}
	fmt.Println(minimalKSum(arr1, 6))

}
