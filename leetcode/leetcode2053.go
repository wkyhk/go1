package leetcode

import "fmt"

/*
2053. 数组中第 K 个独一无二的字符串
独一无二的字符串指的是在一个数组中只出现过 一次的字符串。
给你一个字符串数组arr和一个整数k，请你返回arr中第k个独一无二的字符串。如果少于k个独一无二的字符串，那么返回空字符串""。
注意，按照字符串在原数组中的 顺序找到第 k个独一无二字符串。
链接：https://leetcode.cn/problems/kth-distinct-string-in-an-array
*/
func kthDistinct(arr []string, k int) string {
	m := make(map[string]int, 0)
	for _, v := range arr {
		m[v]++
	}
	i := 0
	var ans string
	for _, v := range arr {
		if m[v] == 1 {
			i++
		}
		if i == k {
			ans = v
			break
		}
	}
	return ans
}
func Example2053() {
	arr := []string{"b", "a", "c", "a"}
	fmt.Println(kthDistinct(arr, 2))
}
