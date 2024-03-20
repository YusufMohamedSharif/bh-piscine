package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

// function to insert element in the last position of the list
func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}

	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}

// type node struct {
// 	data int
// 	next *node
// }

// func insert(head *node, data int) *node {
// 	n := &node{data: data}

// 	if head == nil {
// 		return n
// 	} else {
// 		n.next = head
// 		return n
// 	}
// }

// func PrintList(head *node) {
// 	for head != nil {
// 		fmt.Print(head.data, " -> ")
// 		head = head.next
// 	}
// 	fmt.Println(nil)
// }

// func main() {
// 	var link *node

// 	link = insert(link, 1)
// 	link = insert(link, 2)
// 	link = insert(link, 3)

// 	PrintList(link)
// }
