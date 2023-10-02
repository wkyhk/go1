package listnode

/*
234. 回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false
链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9
*/
func IsPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for ; head != nil; head = head.Next {
		nums = append(nums, head.Val)
	}
	return isPalind(nums)
}
func isPalind(nums []int) bool {
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}
	return true
}
