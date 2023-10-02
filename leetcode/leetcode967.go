package leetcode

import (
	"fmt"
	"math"
)

/*
967. 连续差相同的数字
返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。
请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0 是有效的。
你可以按 任何顺序 返回答案。
链接：https://leetcode.cn/problems/numbers-with-same-consecutive-differences
*/
func numsSameConsecDiff(n int, k int) []int {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 1; i < n; i++ { // n位
		var res []int
		for _, v := range s {
			d := v % 10 //尾数

			if math.Abs(float64(d-k)) < 10 && int(math.Abs(float64(d-k))) != d+k && d >= k {
				res = append(res, v*10+int(math.Abs(float64(d-k))))
			}
			if d+k < 10 {
				res = append(res, v*10+(d+k))
			}

		}

		s = res
	}

	return s

}
func TestLeetCode967() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(2, 0))
	fmt.Println(numsSameConsecDiff(2, 2))
}
