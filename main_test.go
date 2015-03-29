package main

import (
    "testing"
)

/* 
      0
     / \
    1   2
   /
  4
 / \
3   5
*/

func TestExtractNodes(t *testing.T) {
    in := "INSERT INTO `foo` VALUES (0,'name 0',null),(1,'name 1',0),(2,'name 2',0),(3,'name 3',4),(4,'name 4',1),(5,'name 5',4)"

    expectedNodes := []RawAdjacencyTreeNode{
        RawAdjacencyTreeNode{Id: "0", ParentId: "null"},
        RawAdjacencyTreeNode{Id: "1", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "2", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "3", ParentId: "4"},
        RawAdjacencyTreeNode{Id: "4", ParentId: "1"},
        RawAdjacencyTreeNode{Id: "5", ParentId: "4"},
    }

    expectedOut := RawAdjacencyTreeNodes{Nodes: expectedNodes}

    actualOut := extractNodes(in)

    if !actualOut.equalTo(expectedOut) {
        t.Errorf("Expected %v, got %v", expectedOut, actualOut)
    }
}

func TestBuildLinkedNodes(t *testing.T) {
    in := []RawAdjacencyTreeNode{
        RawAdjacencyTreeNode{Id: "0", ParentId: "null"},
        RawAdjacencyTreeNode{Id: "1", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "2", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "3", ParentId: "4"},
        RawAdjacencyTreeNode{Id: "4", ParentId: "1"},
        RawAdjacencyTreeNode{Id: "5", ParentId: "4"},
    }

    elem3 := LinkedAdjacencyTreeNode{Id: "3"}
    elem5 := LinkedAdjacencyTreeNode{Id: "5"}
    elem4 := LinkedAdjacencyTreeNode{Id: "4", Children: []*LinkedAdjacencyTreeNode{&elem3, &elem5}}
    
    elem1 := LinkedAdjacencyTreeNode{Id: "1", Children: []*LinkedAdjacencyTreeNode{&elem4}}
    elem2 := LinkedAdjacencyTreeNode{Id: "2"}

    elem0 := LinkedAdjacencyTreeNode{Id: "0", Children: []*LinkedAdjacencyTreeNode{&elem1, &elem2}}

    expectedOut := elem0

    actualOut := buildLinkedNodes(in)

    if !actualOut.equalTo(expectedOut) {
        t.Errorf("Expected %v, got %v", expectedOut, actualOut)
    }
}

/* 
        1-0-12
       /     \
    2-1-9  10-2-11
     /
   3-4-8
   /   \
4-3-5 6-5-7
*/
func TestAttachedLeftsAndRights(t *testing.T) {
    inElem3 := LinkedAdjacencyTreeNode{Id: "3"}
    inElem5 := LinkedAdjacencyTreeNode{Id: "5"}
    inElem4 := LinkedAdjacencyTreeNode{Id: "4", Children: []*LinkedAdjacencyTreeNode{&inElem3, &inElem5}}
    
    inElem1 := LinkedAdjacencyTreeNode{Id: "1", Children: []*LinkedAdjacencyTreeNode{&inElem4}}
    inElem2 := LinkedAdjacencyTreeNode{Id: "2"}

    inElem0 := LinkedAdjacencyTreeNode{Id: "0", Children: []*LinkedAdjacencyTreeNode{&inElem1, &inElem2}}

    in := inElem0

    outElem3 := LinkedAdjacencyTreeNode{Id: "3", Left: 4, Right: 5}
    outElem5 := LinkedAdjacencyTreeNode{Id: "5", Left: 6, Right: 7}
    outElem4 := LinkedAdjacencyTreeNode{Id: "4", Left: 3, Right: 8, Children: []*LinkedAdjacencyTreeNode{&outElem3, &outElem5}}
    
    outElem1 := LinkedAdjacencyTreeNode{Id: "1", Left: 2, Right: 9, Children: []*LinkedAdjacencyTreeNode{&outElem4}}
    outElem2 := LinkedAdjacencyTreeNode{Id: "2", Left: 10, Right: 11}

    outElem0 := LinkedAdjacencyTreeNode{Id: "0", Left: 1, Right: 12, Children: []*LinkedAdjacencyTreeNode{&outElem1, &outElem2}}

    expectedOut := outElem0

    attachLeftsAndRights(&in)

    if !in.equalTo(expectedOut) {
        t.Errorf("Expected %v, got %v", expectedOut, in)
    }
}
