package leetcode

/*
1652. 拆炸弹
你有一个炸弹需要拆除，时间紧迫！你的情报员会给你一个长度为 n 的 循环 数组 code 以及一个密钥 k 。

为了获得正确的密码，你需要替换掉每一个数字。所有数字会 同时 被替换。

如果 k > 0 ，将第 i 个数字用 接下来 k 个数字之和替换。
如果 k < 0 ，将第 i 个数字用 之前 k 个数字之和替换。
如果 k == 0 ，将第 i 个数字用 0 替换。
由于 code 是循环的， code[n-1] 下一个元素是 code[0] ，且 code[0] 前一个元素是 code[n-1] 。

给你 循环 数组 code 和整数密钥 k ，请你返回解密后的结果来拆除炸弹！

链接:https://leetcode.cn/problems/defuse-the-bomb/description/
*/
func decrypt(code []int, k int) []int {

	ans := make([]int, len(code))
	if k == 0 {
		return ans
	}
	for i, _ := range code {
		ans[i] = sumk(code, i, k)
	}
	return ans
}
func sumk(code []int, index, k int) int {
	sum := 0
	if k > 0 {
		for i := 0; i < k; i++ {
			index = index + 1
			if index >= len(code) {
				index = 0
			}
			sum += code[index]
		}
	} else if k < 0 {
		k = -k
		for i := 0; i < k; i++ {
			index = index - 1
			if index < 0 {
				index = len(code) - 1
			}
			sum += code[index]
		}
	}
	return sum
}
