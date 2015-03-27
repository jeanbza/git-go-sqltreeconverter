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

type LinkedAdjacencyTreeNode struct {
    Id string
    Parent *LinkedAdjacencyTreeNode
}

type LinkedAdjacencyTreeNodes struct {
    Nodes []LinkedAdjacencyTreeNode
}

func (a LinkedAdjacencyTreeNodes) equalTo(b LinkedAdjacencyTreeNodes) bool {
    if len(a.Nodes) != len(b.Nodes) {
        return false
    }

    for index, elemA := range a.Nodes {
        elemB := b.Nodes[index]

        if elemA.Id != elemB.Id {
            return false
        }

        if elemA.Parent != nil && elemA.Parent.Id != elemB.Parent.Id {
            return false
        }
    }

    return true
}

func (a LinkedAdjacencyTreeNode) String() string {
    var parentId string

    if a.Parent == nil {
        parentId = "nil"
    } else {
        parentId = a.Parent.Id
    }

    return fmt.Sprintf("{id: %s, parentId: %s}", a.Id, parentId)
}