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
    Children []*LinkedAdjacencyTreeNode
}

func (a LinkedAdjacencyTreeNode) equalTo(b LinkedAdjacencyTreeNode) bool {
    if a.Id != b.Id {
        return false
    }

    if len(a.Children) != len(b.Children) {
        return false
    }

    for index, _ := range a.Children {
        // Note: Order is considered
        if !a.Children[index].equalTo(*b.Children[index]) {
            return false
        }
    }

    return true
}

func (a LinkedAdjacencyTreeNode) String() string {
    var childrenIds []string
    
    for _, child := range a.Children {
        childrenIds = append(childrenIds, child.String())
    }

    return fmt.Sprintf("{id: %s, children: %s}", a.Id, childrenIds)
}