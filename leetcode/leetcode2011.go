package leetcode

import (
	"fmt"
	"strings"
)

/*
2011. 执行操作后的变量值
存在一种仅支持 4 种操作和 1 个变量 X 的编程语言：
++X 和 X++ 使变量 X 的值 加 1
--X 和 X-- 使变量 X 的值 减 1
最初，X 的值是 0
链接：https://leetcode.cn/problems/final-value-of-variable-after-performing-operations
*/
func finalValueAfterOperations(operations []string) int {
	ans := 0
	op := []string{"++X", "X++", "X--", "--X"}
	for _, v := range operations {
		if strings.Compare(op[0], v) == 0 || strings.Compare(op[1], v) == 0 {
			ans++
		} else if strings.Compare(op[2], v) == 0 || strings.Compare(op[3], v) == 0 {
			ans--
		}
	}
	return ans
}
func Testleetcode2011() {
	operations := []string{"--X", "X++", "X++"}
	fmt.Println(finalValueAfterOperations(operations))
}
