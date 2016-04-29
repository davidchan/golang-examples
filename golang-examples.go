package main

import "fmt"

type Node struct {
	value int   // value
	left  *Node // left
	right *Node // right
}

func sum(node *Node) int {
	left := 0
	right := 0
	if node.left != nil {
		left = sum(node.left)
	}
	if node.right != nil {
		right = sum(node.right)
	}
	return node.value + left + right
}

func sum2(node *Node) int {
	var queue []*Node
	queue = append(queue, node)

	var result = 0
	var current *Node
	for len(queue) > 0 {
		current = queue[0]
		result += current.value

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}

		queue = queue[1:] // shift
	}
	return result
}

func main() {
	tree := Node{1, &Node{2, nil, nil}, &Node{3, nil, &Node{4, nil, nil}}}
	fmt.Println(sum(&tree))
	fmt.Println(sum2(&tree))
}
