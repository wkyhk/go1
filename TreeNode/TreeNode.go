package TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val      int
	Children []*Node
}

func New() *TreeNode {

	root2 := &TreeNode{3, nil, nil}
	root3 := &TreeNode{2, root2, nil}
	root4 := &TreeNode{1, nil, root3}
	return root4
}
