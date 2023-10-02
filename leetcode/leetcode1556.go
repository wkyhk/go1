package leetcode

import "strconv"

/*
1556. 千位分隔数
给你一个整数 n，请你每隔三位添加点（即 "." 符号）作为千位分隔符，并将结果以字符串格式返回。
*/
func ThousandSeparator(n int) string {
	ans := []byte(strconv.Itoa(n))
	for i := len(ans) - 3; i > 0; i = i - 3 {
		ans = append(ans[:i+1], ans[i:]...)
		ans[i] = '.'
	}
	return string(ans)
}
