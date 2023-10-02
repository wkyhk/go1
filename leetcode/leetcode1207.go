package leetcode

import "sort"

/*
1207. 独一无二的出现次数
给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
*/
func UniqueOccurrences(arr []int) bool {
	numCnt := make(map[int]int)

	for _, v := range arr {
		numCnt[v]++
	}

	cnt := []int{}
	for _, count := range numCnt {
		cnt = append(cnt, count)
	}
	sort.Ints(cnt)

	for i := 1; i < len(cnt); i++ {
		if cnt[i] == cnt[i-1] {
			return false
		}
	}
	return true

}
