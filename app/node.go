package main

import (
	"fmt"
	"strings"
)

func (n *Node) CreateNode(nl NodeList) {
	top := nl[0]
	nextList := append(nl[:0], nl[1:]...)

	if len(n.Children) == 0 {
		if len(nl) > 1 {
			top.CreateNode(nextList)
		}

		n.Children = append(n.Children, top)

		return
	}

	for _, child := range n.Children {
		if child.Name == top.Name {
			if len(nl) > 1 {
				child.CreateNode(nextList)

				return
			} else {
				fmt.Printf("There is already a directory called %s in this directory\n", child.Name)

				return
			}
		}
	}

	n.Children = append(n.Children, top)
}

func (n *Node) DeleteNode(nl NodeList, p Path) {
	top := nl[0]
	nextList := append(nl[:0], nl[1:]...)

	if len(nl) == 1 {
		n.Children = append(n.Children[:0], n.Children[1:]...)
		return
	}

	for i, child := range n.Children {
		if child.Name == top.Name {
			if len(nl) == 1 {
				n.Children = append(n.Children[:i], n.Children[i+1:]...)

				return
			}

			child.DeleteNode(nextList, p)
			return
		}
	}

	fmt.Printf("Cannot delete %s - %s does not exist\n", string(p), top.Name)
}

func (n *Node) ReturnNode(nl NodeList) Node {
	var node = Node{}
	top := nl[0]
	nextList := append(nl[:0], nl[1:]...)

	for i, child := range n.Children {
		if child.Name == top.Name {
			if len(nl) > 1 {
				node = child.ReturnNode(nextList)

				return node
			}

			node = *child
			n.Children = append(n.Children[:i], n.Children[i+1:]...)

			return node
		}
	}

	return node
}

func (n Node) Walk() {
	depthStr := strings.Repeat("    ", n.Depth)
	fmt.Println(depthStr, n.Name)

	for _, child := range n.Children {

		child.Walk()
	}
}

func (n Node) UpdateChildren() {
	for _, child := range n.Children {
		child.Depth = n.Depth + 1

		child.UpdateChildren()
	}
}
