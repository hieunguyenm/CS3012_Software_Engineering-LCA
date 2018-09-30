package lca

// Conor Gildea 2018/2019 - TCD SCSS ICS JS

// First Golang project and CI project
// Used golang.org/x/tour/tree as a basis for insert function
// Converted functions given here from C++ to GoLang
// https://www.geeksforgeeks.org/lowest-common-ancestor-binary-tree-set-1/

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

func findPath(tree *Tree,path *[]int, key int)  bool {
	if(tree == nil){
		return false;
	}
	*path = append(*path,tree.Value)
	if tree.Value == key{
		return true;
	}
	if (findPath(tree.Left,path,key))||
	(findPath(tree.Right,path,key)){
		return true;
	}
	n := len(*path) - 1
	*path =(*path)[:n]
	return false
}

func findLCA(tree *Tree, keyOne int, keyTwo int) int{
	var keyOnePath [] int
	var keyTwoPath [] int

	if(!findPath(tree,&keyOnePath,keyOne)||!findPath(tree,&keyTwoPath,keyTwo)){
		return -1
	}
	var i int
	for i=0;i<len(keyOnePath) && i<len(keyTwoPath);i++{
		if(keyOnePath[i]!=keyTwoPath[i]){
			break
		}
	}
	return keyOnePath[i-1]

}
