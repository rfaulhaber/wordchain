package node

import (
    "os"
    "bufio"
)

type Node struct {
    Edge       string
    Children   []*Node
}

const TERMINAL_EDGE = "*"
const ROOT_EDGE = "$"

func MakeNode(e string) *Node {
    return &Node{e, nil}
}

func (n *Node) AddChild(e string) {
    n.Children = append(n.Children, MakeNode(e))
}

func (n *Node) HasChild(c string) (bool, int) {
    for i, node := range n.Children {
        if string(node.Edge[0]) == c {
            return true, i
        }
    }

    return false, -1;
}

func (n *Node) HasWord(word string) bool {
    if len(word) == 1 {
        return string(n.Edge[0]) == string(word[0]) && string(n.Edge[1]) == TERMINAL_EDGE
    } else {
        // TODO: do I need to make this distinction?
        if string(n.Edge[0]) == ROOT_EDGE {
            if hasChild, index := n.HasChild(string(word[0])); hasChild {
                return n.Children[index].HasWord(word[1:])
            } else {
                return false;
            }
        } else {
            if hasChild, index := n.HasChild(string(word[0])); hasChild {
                return n.HasWord(word[1:])
            } else {
                return false;
            }
        }
    }
}

func BuildTree(inputFilePath string) *Node {
	root := MakeNode(ROOT_EDGE) // special root character

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

		if hasChild, index := n.HasChild(c); hasChild {
            if suffix == "" {
                n.Children[index].Edge += TERMINAL_EDGE
            }

			DescendTreeAddChild(suffix, n.Children[index])
		} else {
            if suffix == "" {
                c += TERMINAL_EDGE
            }
			n.AddChild(c)
			DescendTreeAddChild(suffix, n.Children[len(n.Children) - 1])
		}
	}
}
