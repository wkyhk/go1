/*
2441. 与对应负数同时存在的最大正整数
给你一个 不包含 任何零的整数数组 nums ，找出自身与对应的负数都在数组中存在的最大正整数 k 。

返回正整数 k ，如果不存在这样的整数，返回 -1 。
链接：https://leetcode.cn/problems/largest-positive-integer-that-exists-with-its-negative/?envType=daily-question&envId=2023-10-05
*/
package leetcode

func findMaxK(nums []int) int {
	k := -1
	m := make(map[int]struct{}, 0)
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for i, _ := range m {
		if i > 0 {
			if _, ok := m[-i]; ok {
				k = max(k, i)
			}
		}

	}
	return k
}
