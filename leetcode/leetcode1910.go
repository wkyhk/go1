package leetcode

/*
1910. 删除一个字符串中所有出现的给定子字符串
给你两个字符串s和part，请你对s反复执行以下操作直到 所有子字符串part都被删除：

找到 s中 最左边的子字符串 part，并将它从 s中删除。
请你返回从 s中删除所有 part子字符串以后得到的剩余字符串。

一个 子字符串是一个字符串中连续的字符序列。

链接：https://leetcode.cn/problems/remove-all-occurrences-of-a-substring

*/
func removeOccurrences(s string, part string) string {
	if len(s) == 0 {
		return s
	}
	stack := []byte{}
	sn, pn := len(s), len(part)
	for i := 0; i < sn; i++ {
		stack = append(stack, s[i])
		if s[i] == part[pn-1] {
			if len(stack) >= pn {
				st := string(stack[len(stack)-pn:])
				if st == part {
					stack = stack[:len(stack)-pn]
				}

			}
		}
	}
	return string(stack)
}
