package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val  int
	Next *Node
}

var head = new(Node)

func addNode(t *Node, v int) int {
	if t == nil {
		t = &Node{v, nil}
		return 0
	}

	//if v == t.Val {
	//	fmt.Println("节点已存在：", v)
	//	return -1
	//}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Val)
		t = t.Next
	}

	fmt.Println()
}

func lookupNode(t *Node, v int) bool {
	if head == nil {
		t = &Node{v, nil}
		head = t
		return false
	}

	if v == t.Val {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> 空链表!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	// 暴力解法，循环取数并计算，得到数值
	var l1List []int
	var l2List []int
	var count = 0
	var result1 = 0

	for l1 != nil {
		result1 += l1.Val * int(math.Pow10(count))
		l1List = append(l1List, l1.Val)
		l1 = l1.Next
		count++
	}
	for l2 != nil {
		l2List = append(l2List, l2.Val)
		l2 = l2.Next
	}
	return nil
}

var node1 = new(Node)
var node2 = new(Node)

func main() {
	//node1 = nil
	//node2 = nil
	addNode(node1, 1)
	addNode(node1, 2)
	addNode(node1, 3)
	addNode(node2, 3)
	addNode(node2, 3)
	addNode(node2, 3)

	traverse(node1)

	traverse(node2)
}
