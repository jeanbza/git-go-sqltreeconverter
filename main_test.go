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

    expectedOut := []RawAdjacencyTreeNode{
        RawAdjacencyTreeNode{Id: "0", ParentId: "null"},
        RawAdjacencyTreeNode{Id: "1", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "2", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "3", ParentId: "4"},
        RawAdjacencyTreeNode{Id: "4", ParentId: "1"},
        RawAdjacencyTreeNode{Id: "5", ParentId: "4"},
    }

    actualOut := extractNodes(in)

    if !rawAdjacencyTreeNodeMatches(actualOut, expectedOut) {
        t.Errorf("Expected %v, got %v", expectedOut, actualOut)
    }
}

func rawAdjacencyTreeNodeMatches(a, b []RawAdjacencyTreeNode) bool {
    if len(a) != len(b) {
        return false
    }

    for index, elemA := range a {
        elemB := b[index]

        if elemA != elemB {
            return false
        }
    }

    return true
}