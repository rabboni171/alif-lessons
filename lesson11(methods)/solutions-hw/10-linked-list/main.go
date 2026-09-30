package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (head *Node) Append(v int) *Node {
	newNode := &Node{Value: v}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return head
}

func PrintList(head *Node) {
	for n := head; n != nil; n = n.Next {
		if n.Next != nil {
			fmt.Printf("%d -> ", n.Value)
		} else {
			fmt.Println(n.Value)
		}
	}
}

func main() {
	head := &Node{Value: 1}
	head.Append(2)
	head.Append(3)
	head.Append(4)
	head.Append(5)

	PrintList(head)
}
