package datastructures

// BinaryTree is next
type BinaryTree struct {
	Root *TreeNode
}

// TreeNode is ok i guess
type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

// InsertLeft of node
func (tn *TreeNode) InsertLeft(value int) {

	var newNode TreeNode
	newNode.Value = value

	if tn.Left == nil {
		tn.Left = &newNode
	} else {
		newNode.Left = tn.Left
		tn.Left = &newNode
	}
}

// InsertRight of node
func (tn *TreeNode) InsertRight(value int) {

	var newNode TreeNode
	newNode.Value = value

	if tn.Right == nil {
		tn.Right = &newNode
	} else {
		newNode.Right = tn.Right
		tn.Right = &newNode
	}
}
