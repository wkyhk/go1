package listnode

func MergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	preHead := &ListNode{Val: -1} // 虚拟头结点
	pre := preHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			pre.Next = head1
			head1 = head1.Next
			pre = pre.Next
		} else {
			pre.Next = head2
			head2 = head2.Next
			pre = pre.Next
		}
	}
	if head1 == nil {
		pre.Next = head2
	}
	if head2 == nil {
		pre.Next = head1
	}
	return preHead.Next

}
