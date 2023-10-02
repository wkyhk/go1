package leetcode

import "fmt"

/*
946. 验证栈序列
给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。
链接：https://leetcode.cn/problems/validate-stack-sequences
*/
func validateStackSequences(pushed []int, popped []int) bool {
	st := make([]int, 0, len(pushed))
	curP := 0
	for i := 0; i < len(pushed); i++ {
		st = append(st, pushed[i])
		for j := len(st) - 1; j >= 0; j-- {
			if st[j] == popped[curP] {
				st = st[:j]
				curP++
			} else {
				break
			}
		}
	}
	return len(st) == 0

}
func Example964() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}

	pushed1 := []int{1, 2, 3, 4, 5}
	popped1 := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed, popped))
	fmt.Println(validateStackSequences(pushed1, popped1))

}
