package piscine

type TreeNodee struct {
	Data  string
	Left  *TreeNodee
	Right *TreeNodee
}

func printOrderTraversal(root *TreeNodee, niveau int, fn func(string)) {
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

func BTreeApplyByLevel(root *TreeNodee, fn func(string)) {
	h := BTreeLevelCount(root)
	for i := 1; i <= h; i++ {
		printOrderTraversal(root, i, fn)
	}
}

func BTreeLevelCount(root *TreeNodee) int {
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
