package listnode

import "sort"

/*
148. 排序链表
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/
func SortList(head *ListNode) *ListNode {
	nums := make([]int, 0)
	h := head
	for h != nil {
		nums = append(nums, h.Val)
		h = h.Next
	}
	sort.Ints(nums)
	h = head
	for i := 0; i < len(nums); i++ {
		h.Val = nums[i]
		h = h.Next
	}
	return head
}
