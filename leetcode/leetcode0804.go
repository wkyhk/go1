package leetcode

import "fmt"

/*
面试题 08.04. 幂集
幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
说明：解集不能包含重复的子集。
https://leetcode.cn/problems/power-set-lcci/
*/
func subsets1(nums []int) [][]int {
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

func Example0804() {
	arr := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(arr))
}
