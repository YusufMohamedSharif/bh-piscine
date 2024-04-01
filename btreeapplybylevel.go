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

// Define a wrapper function that takes a string argument and calls the original function f
func wrapperFn(data string, f func(...interface{}) (int, error)) {
	// Ignore the string argument and call the original function f
	f()
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
	// Pass the wrapper function to printOrderTraversal
	for i := 1; i <= h; i++ {
		printOrderTraversal(root, i, func(data string) {
			wrapperFn(data, f)
		})
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
