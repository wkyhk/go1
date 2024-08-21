package leetcode

import "math/bits"

/*
3007. 价值和小于等于 K 的最大数字
给你一个整数 k 和一个整数 x 。整数 num 的价值是它的二进制表示中在 x，2x，3x 等位置处
设置位
 的数目（从最低有效位开始）。下面的表格包含了如何计算价值的例子。
 num 的 累加价值 是从 1 到 num 的数字的 总 价值。如果 num 的累加价值小于或等于 k 则被认为是 廉价 的。

请你返回 最大 的廉价数字。
链接：https://leetcode.cn/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/description/
*/
func findMaximumNumber(k int64, x int) int64 {
	l, r := int64(1), (k+1)<<x
	for l < r {
		m := (l + r + 1) / 2
		if accumulatedPrice(x, m) > k {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func accumulatedBitPrice(x int, num int64) int64 {
	period := int64(1) << x
	res := period / 2 * (num / period)
	if num%period >= period/2 {
		res += num%period - (period/2 - 1)
	}
	return res
}

func accumulatedPrice(x int, num int64) int64 {
	res := int64(0)
	length := 64 - bits.LeadingZeros64(uint64(num))
	for i := x; i <= length; i += x {
		res += accumulatedBitPrice(i, num)
	}
	return res
}
