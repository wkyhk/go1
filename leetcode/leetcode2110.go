package leetcode

/*
2110. 股票平滑下跌阶段的数目
给你一个整数数组prices，表示一支股票的历史每日股价，其中prices[i]是这支股票第i天的价格。
一个 平滑下降的阶段定义为：对于连续一天或者多天，每日股价都比 前一日股价恰好少 1，这个阶段第一天的股价没有限制。
请你返回 平滑下降阶段的数目。
链接：https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock
*/
func getDescentPeriods(prices []int) int64 {
	var ans int64 = 0
	dp := 1
	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i-1] == prices[i]+1 {
			dp++
		} else {
			dp = 1
		}
		ans += int64(dp)
	}
	return ans
}
