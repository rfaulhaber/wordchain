package node

import (
    "os"
    "bufio"
)

type Node struct {
    Edge string
    Children []*Node
}

func MakeNode(e string) *Node {
    return &Node{e, nil}
}

func (n *Node) AddChild(e string) {
    n.Children = append(n.Children, MakeNode(e))
}

func (n *Node) HasChild(c string) (bool, int) {
    for i, node := range n.Children {
        if node.Edge == c {
            return true, i;
        }
    }

    return false, -1;
}

func BuildTree(inputFilePath string) *Node {
	root := MakeNode("$") // special root character

	f, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f);
	for scanner.Scan() {
        word := scanner.Text()
		DescendTreeAddChild(word, root)
	}

    return root
}

// recursively adds a node to a tree
func DescendTreeAddChild(word string, n *Node) {
	if len(word) > 0 {
		c := string(word[0])
		suffix := word[1:]

		if b, index := n.HasChild(c); b {
			DescendTreeAddChild(suffix, n.Children[index])
		} else {
			n.AddChild(c)
			DescendTreeAddChild(suffix, n.Children[len(n.Children) - 1])
		}
	}
}
