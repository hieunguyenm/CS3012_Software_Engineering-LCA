package lca

// Conor Gildea 2018/2019 - TCD SCSS ICS JS

// First Golang project and CI project
// Used golang.org/x/tour/tree as a basis for insert function


type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func insert(tree *Tree, value int) *Tree {
	if tree == nil {
		return &Tree{nil, value, nil}
	}
	if value < tree.Value {
		tree.Left = insert(tree.Left, value)
	} else {
		tree.Right = insert(tree.Right, value)
	}
	return tree
}

func Hello() string {
	return "Hello";
}
