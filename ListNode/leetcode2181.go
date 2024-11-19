package listnode

/*
2181. 合并零之间的节点
给你一个链表的头节点 head ，该链表包含由 0 分隔开的一连串整数。链表的 开端 和 末尾 的节点都满足 Node.val == 0 。

对于每两个相邻的 0 ，请你将它们之间的所有节点合并成一个节点，其值是所有已合并节点的值之和。然后将所有 0 移除，修改后的链表不应该含有任何 0 。

	返回修改后链表的头节点 head 。

链接：https://leetcode.cn/problems/merge-nodes-in-between-zeros/description/
*/
func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	total := 0
	for cur := head.Next; cur != nil; cur = cur.Next {
		if cur.Val == 0 {
			node := &ListNode{Val: total}
			tail.Next = node
			tail = tail.Next
			total = 0
		} else {
			total += cur.Val
		}
	}
	return dummy.Next
}