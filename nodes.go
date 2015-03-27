package main

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

        if elemA != elemB {
            return false
        }
    }

    return true
}