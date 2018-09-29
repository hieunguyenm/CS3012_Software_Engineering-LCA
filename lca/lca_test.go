package lca

import "testing"

func Test_Hello(t *testing.T) {
  var v string
  v = Hello()
  if v != "Hello" {
    t.Error("Expected \"Hello\", got ", v)
  }
}

func Test_LCA_Root_Is_LCA(t *testing.T) {
	tree:= &node{
		value:4
	}
	tree = insertFromRoot(tree,1)
    result = findLCA(tree,1)
	if result != 4 {
    t.Error("Expected \"4\", got ", v)
  }
}

func Test_LCA_Node_Not_Exist(t *testing.T) {
	tree:= &node{
		value:4
	}
    result = findLCA(tree,1)
	if result != nil {
    t.Error("Expected \"nil\" as node with value 1 doesn't exist got ", v)
  }
}

func Test_LCA_Good_Behvaiour_Case(t *testing.T) {
	tree:= &node{
		value:4
	}
	tree = insertFromRoot(tree,1)
    tree = insertFromRoot(tree,2)
	result = findLCA(1,2)
	if result != 3 {
    t.Error("Expected \"4\", got ", v)
  }
}
