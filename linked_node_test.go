package main

import (
    "testing"
)

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

    in.attachLeftsAndRights()

    if !in.equalTo(expectedOut) {
        t.Errorf("Expected %v, got %v", expectedOut, in)
    }
}
