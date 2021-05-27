package main

import "LeanGo/tree"

func main() {
	var root tree.TreetreeNode
	root = treeNode{
		value: 3,
	}
	root.left = &treeNode{}
	root.right = &treeNode{5, nill, nill}
	root.right.right = new(treeNode)

}
