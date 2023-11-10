package leetcode

import "sort"

/*
2300. 咒语和药水的成功对数
给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[j] 表示第 j 瓶药水的能量强度。
同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合。
请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目。

示例 1：
输入：spells = [5,1,3], potions = [1,2,3,4,5], success = 7
输出：[4,1,3]
解释：
- 咒语 1（5 x 1）和药水 1（1 x 1）的组合成功。
- 咒语 1（5 x 1）和药水 2（1 x 2）的组合成功。
- 咒语 1（5 x 1）和药水 3（1 x 3）的组合成功。
- 咒语 2（1 x 3）和药水 3（1 x 3）的组合成功。
- 咒语 3（3 x 3）和药水 3（1 x 3）的组合成功。
总共 4 个成功组合，所以返回 [4,1,3] 。
链接：https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/description/
*/
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := make([]int, len(spells))
	for i, v := range spells {
		ans[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/v+1)
	}
	return ans
}
