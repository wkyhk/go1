package listnode

/*
2487. 从链表中移除节点
给你一个链表的头节点 head 。
移除每个右侧有一个更大数值的节点。
返回修改后链表的头节点 head 。
链接：https://leetcode.cn/problems/remove-nodes-from-linked-list/?envType=daily-question&envId=2024-02-23
*/
func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := make([]*ListNode, 0)
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	max := res[len(res)-1].Val
	for i := len(res) - 1; i >= 0; i-- {
		if res[i].Val >= max {
			max = res[i].Val
		} else {
			res[i].Next = nil
			res = append(res[:i], res[i+1:]...)
		}
	}
	for i := 0; i < len(res)-1; i++ {
		res[i].Next = res[i+1]
	}
	res[len(res)-1].Next = nil
	return res[0]
}
