package main

import (
    "testing"
)

// 0
// 1 2
// 4
// 3 5

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

// func TestBuildLinkedNodes(t *testing.T) {
//     in := []RawAdjacencyTreeNode{
//         RawAdjacencyTreeNode{Id: "0", ParentId: "null"},
//         RawAdjacencyTreeNode{Id: "1", ParentId: "0"},
//         RawAdjacencyTreeNode{Id: "2", ParentId: "0"},
//         RawAdjacencyTreeNode{Id: "3", ParentId: "4"},
//         RawAdjacencyTreeNode{Id: "4", ParentId: "1"},
//         RawAdjacencyTreeNode{Id: "5", ParentId: "4"},
//     }

//     elem0 := *LinkedAdjacencyTreeNode{Id: "0"}
//     elem1 := *LinkedAdjacencyTreeNode{Id: "1", Parent: elem0}
//     elem2 := *LinkedAdjacencyTreeNode{Id: "2", Parent: elem0}
//     elem4 := *LinkedAdjacencyTreeNode{Id: "4", Parent: elem1}
//     elem3 := *LinkedAdjacencyTreeNode{Id: "3", Parent: elem4}
//     elem5 := *LinkedAdjacencyTreeNode{Id: "5", Parent: elem4}


// }