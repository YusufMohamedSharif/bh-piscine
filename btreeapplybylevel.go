package piscine

type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func printOrderTraversal(root *TreeNode, niveau int, fn func(string)) {
	if root == nil {
		return
	}
	if niveau == 1 {
		fn(root.Data)
	} else if niveau > 1 {
		printOrderTraversal(root.Left, niveau-1, fn)
		printOrderTraversal(root.Right, niveau-1, fn)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 1; i <= h; i++ {
		// printOrderTraversal(root, i, f)
	}
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lHeight := BTreeLevelCount(root.Left)
	rHeight := BTreeLevelCount(root.Right)

	if lHeight > rHeight {
		return lHeight + 1
	}
	return rHeight + 1
}
