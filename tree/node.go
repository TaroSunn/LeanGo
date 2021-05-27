package tree

type treeNode struct {
	value int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) SetValue(value int) {
	node.value = value
}

