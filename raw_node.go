package main

import (
    "fmt"
)

type RawAdjacencyTreeNode struct {
    Id, ParentId string
}

type RawAdjacencyTreeNodes struct {
    Nodes []RawAdjacencyTreeNode
}

func (a RawAdjacencyTreeNode) String() string {
    return fmt.Sprintf("{Id: %s, ParentId: %s}", a.Id, a.ParentId)
}

func (a RawAdjacencyTreeNodes) equalTo(b RawAdjacencyTreeNodes) bool {
    if len(a.Nodes) != len(b.Nodes) {
        return false
    }

    for index, elemA := range a.Nodes {
        elemB := b.Nodes[index]

        if elemA != elemB {
            return false
        }
    }

    return true
}

func (a RawAdjacencyTreeNode) isRoot() bool {
    return a.ParentId == "null" || a.Id == a.ParentId
}