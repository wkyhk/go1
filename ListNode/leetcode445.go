package listnode

/*
445. 两数相加 II
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。
链接：https://leetcode.cn/problems/add-two-numbers-ii
*/
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := make([]*ListNode, 0)
	list2 := make([]*ListNode, 0)
	for l1 != nil {
		list1 = append(list1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		list2 = append(list2, l2)
		l2 = l2.Next
	}
	if len(list2) > len(list1) {
		list1, list2 = list2, list1
	}
	temp := 0
	for i, j := len(list1)-1, len(list2)-1; i >= 0; i, j = i-1, j-1 {
		if j >= 0 {
			list1[i].Val = list1[i].Val + list2[j].Val + temp

		} else {
			list1[i].Val = list1[i].Val + temp
		}
		temp = 0
		if list1[i].Val >= 10 {
			temp = list1[i].Val / 10
			list1[i].Val = list1[i].Val % 10
		}
	}
	if temp != 0 {
		return &ListNode{Val: temp, Next: list1[0]}
	}
	return list1[0]
}
