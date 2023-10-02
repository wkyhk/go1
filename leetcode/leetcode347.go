package leetcode

/*
347. 前 K 个高频元素
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
https://leetcode.cn/problems/top-k-frequent-elements/
*/
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	buckets := make([][]int, len(nums)+1)
	for num, v := range m {
		if len(buckets[v]) <= 0 {
			buckets[v] = make([]int, 0)
		}
		buckets[v] = append(buckets[v], num)
	}
	ret := make([]int, 0)
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) > 0 {
			ret = append(ret, buckets[i]...)
			if len(ret) >= k {
				return ret[:k]
			}
		}
	}
	return ret
}
