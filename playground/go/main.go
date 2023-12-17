package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func bfs(root *node) []int {
	if root == nil {
		return []int{}
	}

	queue := []*node{root}
	result := make([]int, 0)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.val)

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return result
}

func main() {
	root := &node{val: 1}
	root.left = &node{val: 2}
	root.right = &node{val: 3}
	root.left.left = &node{val: 4}
	root.left.right = &node{val: 5}
	root.right.left = &node{val: 6}
	root.right.right = &node{val: 7}

	fmt.Println(bfs(root))
}
