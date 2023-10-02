package leetcode

/*
1640. 能否连接形成数组
给你一个整数数组 arr ，数组中的每个整数 互不相同 。另有一个由整数数组构成的数组 pieces，其中的整数也 互不相同 。请你以 任意顺序 连接 pieces 中的数组以形成 arr 。但是，不允许 对每个数组 pieces[i] 中的整数重新排序。
如果可以连接 pieces 中的数组形成 arr ，返回 true ；否则，返回 false
链接：https://leetcode.cn/problemscheck-array-formation-through-concatenation
*/
func canFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int]int, 0)
	for i, v := range pieces {
		m[v[0]] = i
	}
	for i := 0; i < len(arr); {
		j, ok := m[arr[i]]
		if !ok {
			return false
		}
		for _, x := range pieces[j] {
			if arr[i] != x {
				return false
			}
			i++
		}
	}

	return true
}
