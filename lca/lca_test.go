package lca

import "testing"

func Test_LCA_Root_Is_LCA(t *testing.T) {
	root:= insert(nil,4)
	insert(root,1)
	var result int = findLCA(root,4,1)
	if result != 4 {
    t.Error("Expected \"4\", got ", result)
  }
}

func Test_LCA_Root_Is_LCA_Keys_Children_Of_Root(t *testing.T) {
	root:= insert(nil,50)
	insert(root,25)
	insert(root,75)
	var result int = findLCA(root,25,75)
	if result != 50 {
    t.Error("Expected \"50\", got ", result)
  }
}

func Test_LCA_Not_Present_Keys_Children_Of_Root(t *testing.T) {
	root:= insert(nil,50)
	insert(root,25)
	insert(root,75)
	var result int = findLCA(root,25,76)
	if result != -1 {
    t.Error("Expected \"-1\", got ", result)
  }
}


