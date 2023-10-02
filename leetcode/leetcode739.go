package leetcode

/*
739. 每日温度
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
链接：https://leetcode.cn/problems/daily-temperatures
*/
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	st := make([]int, 0)
	for i, v := range temperatures {
		for len(st) > 0 && v > temperatures[st[len(st)-1]] {
			top := st[len(st)-1]
			ans[top] = i - top
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
