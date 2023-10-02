package listnode

/*
面试题 02.06. 回文链表
编写一个函数，检查输入的链表是否是回文的。
https://leetcode.cn/problems/palindrome-linked-list-lcci/
*/
func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}

	return true
}
