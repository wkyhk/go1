package leetcode

import "sort"

/*
881. 救生艇
给定数组 people 。people[i]表示第 i 个人的体重 ，船的数量不限，每艘船可以承载的最大重量为 limit。

每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。

返回 承载所有人所需的最小船数 。
链接：https://leetcode.cn/problems/boats-to-save-people/description/
*/
func numRescueBoats(people []int, limit int) int {
	ans := 0
	sort.Ints(people)
	i, j := 0, len(people)-1
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
		}
		j--
		ans++
	}
	return ans
}
