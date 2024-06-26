package leetcode

import "slices"

/*
2810. 故障键盘
你的笔记本键盘存在故障，每当你在上面输入字符 'i' 时，它会反转你所写的字符串。而输入其他字符则可以正常工作。

给你一个下标从 0 开始的字符串 s ，请你用故障键盘依次输入每个字符。

返回最终笔记本屏幕上输出的字符串。
链接：https://leetcode.cn/problems/faulty-keyboard/description/
*/
func finalString(s string) string {
	str1 := []byte(s)
	ans := make([]byte, 0)

	for i := 0; i < len(str1); i++ {
		if str1[i] == 'i' {
			slices.Reverse(ans)
		} else {
			ans = append(ans, str1[i])
		}
	}
	return string(ans)
}
