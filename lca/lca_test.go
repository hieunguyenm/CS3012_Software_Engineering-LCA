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


func Test_Insert_Correct_Location_Left(t *testing.T) {
	root:= insert(nil,50)
	insert(root,25)
	insert(root,75)
	var result int = root.Left.Value;
	if result != 25 {
    t.Error("Expected \"25\", got ", result)
  }
}

func Test_Insert_Correct_Location_Right(t *testing.T) {
	root:= insert(nil,50)
	insert(root,25)
	insert(root,75)
	var result int = root.Right.Value;
	if result != 75 {
    t.Error("Expected \"75\", got ", result)
  }
}

func Test_Insert_Child_Not_Present(t *testing.T) {
	root:= insert(nil,50)
	insert(root,25)
	var result = root.Right
	if result != nil {
    t.Error("Expected \"nil\", got ", result)
  }
}
