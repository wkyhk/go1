package leetcode

import "strings"

/*
2864. 最大二进制奇数
给你一个 二进制 字符串 s ，其中至少包含一个 '1' 。

你必须按某种方式 重新排列 字符串中的位，使得到的二进制数字是可以由该组合生成的 最大二进制奇数 。

以字符串形式，表示并返回可以由给定组合生成的最大二进制奇数。

注意 返回的结果字符串 可以 含前导零。
链接：https://leetcode.cn/problems/maximum-odd-binary-number/description/
*/
func maximumOddBinaryNumber(s string) string {
	cnt := strings.Count(s, "1")
	return strings.Repeat("1", cnt-1) + strings.Repeat("0", len(s)-cnt) + "1"
}
