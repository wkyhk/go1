package listnode
/*
83. 删除排序链表中的重复元素
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head ==nil{
		return head
	}
    h:=head
	for ;h.Next!=nil;{
		if h.Val==h.Next.Val{
			h.Next=h.Next.Next
		}else {
			h=h.Next
		}
	}
	return head
}