package piscine

type Function func(string)

func printOrderTraversal(root *TreeNode, niveau int, fn Function) {
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

func BTreeApplyByLevel(root *TreeNode, fn Function) {
	h := BTreeLevelCount(root)
	for i := 1; i <= h; i++ {
		printOrderTraversal(root, i, fn)
	}
}
