package listnode

/*
1669. 合并两个链表
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。

请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。

请你返回结果链表的头指针。
链接：https://leetcode.cn/problems/merge-in-between-linked-lists/
*/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	preA := list1
	for i := 0; i < a-1; i++ {
		preA = preA.Next
	}
	preB := preA
	for i := 0; i < b-a+2; i++ {
		preB = preB.Next
	}
	preA.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = preB
	return list1
}
