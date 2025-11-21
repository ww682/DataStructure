package main

import "fmt"

// 定义链表节点结构体
type Node struct {
	Value int
	Next  *Node
}

// 删除链表中的最后一个节点
func DeleteLastNode(head *Node) *Node {
	// 如果链表为空，直接返回 nil
	if head == nil {
		return nil
	}

	// 如果链表只有一个节点，删除它并返回 nil
	if head.Next == nil {
		return nil
	}

	// 遍历链表，找到倒数第二个节点
	current := head
	for current.Next != nil && current.Next.Next != nil {
		current = current.Next
	}

	// 删除最后一个节点
	current.Next = nil
	return head
}

// 打印链表
func PrintList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// 创建一个简单的链表 1 -> 2 -> 3 -> 4 -> nil
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 4}

	fmt.Println("原始链表:")
	PrintList(head)

	// 删除最后一个节点
	head = DeleteLastNode(head)
	fmt.Println("删除最后一个节点后的链表:")
	PrintList(head)
}
