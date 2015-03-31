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

    in := []LinkedAdjacencyTreeNode{inElem0}

    outElem3 := LinkedAdjacencyTreeNode{Id: "3", Left: 4, Right: 5}
    outElem5 := LinkedAdjacencyTreeNode{Id: "5", Left: 6, Right: 7}
    outElem4 := LinkedAdjacencyTreeNode{Id: "4", Left: 3, Right: 8, Children: []*LinkedAdjacencyTreeNode{&outElem3, &outElem5}}
    
    outElem1 := LinkedAdjacencyTreeNode{Id: "1", Left: 2, Right: 9, Children: []*LinkedAdjacencyTreeNode{&outElem4}}
    outElem2 := LinkedAdjacencyTreeNode{Id: "2", Left: 10, Right: 11}

    outElem0 := LinkedAdjacencyTreeNode{Id: "0", Left: 1, Right: 12, Children: []*LinkedAdjacencyTreeNode{&outElem1, &outElem2}}

    expectedOut := []LinkedAdjacencyTreeNode{outElem0}

    actualOut := attachLeftsAndRights(in)

    if len(actualOut) != len(expectedOut) {
        t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
    } else {
        for index := range actualOut {
            if !actualOut[index].equalTo(expectedOut[index]) {
                t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
                break
            }
        }
    }
}

/* 
        1-0-12            13-7-18       19-6-20
       /     \            /     \
    2-1-9  10-2-11    14-4-15 16-8-17
     /
   3-9-8
   /   \
4-3-5 6-5-7
*/
func TestAttachedLeftsAndRights_MultipleRoots(t *testing.T) {
    
    // ------

    inElem3 := LinkedAdjacencyTreeNode{Id: "3"}
    inElem5 := LinkedAdjacencyTreeNode{Id: "5"}
    inElem9 := LinkedAdjacencyTreeNode{Id: "9", Children: []*LinkedAdjacencyTreeNode{&inElem3, &inElem5}}
    
    inElem1 := LinkedAdjacencyTreeNode{Id: "1", Children: []*LinkedAdjacencyTreeNode{&inElem9}}
    inElem2 := LinkedAdjacencyTreeNode{Id: "2"}

    inElem0 := LinkedAdjacencyTreeNode{Id: "0", Children: []*LinkedAdjacencyTreeNode{&inElem1, &inElem2}}

    // --

    inElem4 := LinkedAdjacencyTreeNode{Id: "4"}
    inElem8 := LinkedAdjacencyTreeNode{Id: "8"}

    inElem7 := LinkedAdjacencyTreeNode{Id: "7", Children: []*LinkedAdjacencyTreeNode{&inElem4, &inElem8}}

    // --

    inElem6 := LinkedAdjacencyTreeNode{Id: "6"}

    // ------

    in := []LinkedAdjacencyTreeNode{inElem0, inElem7, inElem6}

    // ------

    outElem3 := LinkedAdjacencyTreeNode{Id: "3", Left: 4, Right: 5}
    outElem5 := LinkedAdjacencyTreeNode{Id: "5", Left: 6, Right: 7}
    outElem9 := LinkedAdjacencyTreeNode{Id: "9", Left: 3, Right: 8, Children: []*LinkedAdjacencyTreeNode{&outElem3, &outElem5}}
    
    outElem1 := LinkedAdjacencyTreeNode{Id: "1", Left: 2, Right: 9, Children: []*LinkedAdjacencyTreeNode{&outElem9}}
    outElem2 := LinkedAdjacencyTreeNode{Id: "2", Left: 10, Right: 11}

    outElem0 := LinkedAdjacencyTreeNode{Id: "0", Left: 1, Right: 12, Children: []*LinkedAdjacencyTreeNode{&outElem1, &outElem2}}

    // --

    outElem4 := LinkedAdjacencyTreeNode{Id: "4", Left: 14, Right: 15}
    outElem8 := LinkedAdjacencyTreeNode{Id: "8", Left: 16, Right: 17}

    outElem7 := LinkedAdjacencyTreeNode{Id: "7", Left: 13, Right: 18, Children: []*LinkedAdjacencyTreeNode{&outElem4, &outElem8}}

    // --

    outElem6 := LinkedAdjacencyTreeNode{Id: "6", Left: 19, Right: 20}

    // ------

    expectedOut := []LinkedAdjacencyTreeNode{outElem0, outElem7, outElem6}

    actualOut := attachLeftsAndRights(in)

    if len(actualOut) != len(expectedOut) {
        t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
    } else {
        for index := range actualOut {
            if !actualOut[index].equalTo(expectedOut[index]) {
                t.Errorf("Problem at index %d.\nIndex expected:\n%v\nIndex got:\n%v\nFull expected:\n%v\nFull got:\n%v", index, expectedOut[index], actualOut[index], expectedOut, actualOut)
                break
            }
        }
    }
}