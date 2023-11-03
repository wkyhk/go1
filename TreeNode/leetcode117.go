package TreeNode

/*
117. 填充每个节点的下一个右侧节点指针 II
给定一个二叉树：

	struct Node {
	  int val;
	  Node *left;
	  Node *right;
	  Node *next;
	}

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

示例 1：

输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
初始状态下，所有 next 指针都被设置为 NULL 。
链接:https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
*/
type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

func connect(root *Node1) *Node1 {
	start := root
	for start != nil {
		var nextStart, last *Node1
		handle := func(cur *Node1) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}
