package main

import (
    "fmt"
)

type LinkedAdjacencyTreeNode struct {
    Id string
    Children []*LinkedAdjacencyTreeNode
    Left, Right int
}

type LinkedAdjacencyTreeNodes struct {
    Nodes []LinkedAdjacencyTreeNode
}

func (a LinkedAdjacencyTreeNode) equalTo(b LinkedAdjacencyTreeNode) bool {
    if a.Id != b.Id || a.Left != b.Left || a.Right != b.Right {
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

    return fmt.Sprintf("{id: %s, left: %d, right: %d, children: %s}", a.Id, a.Left, a.Right, childrenIds)
}

func (root *LinkedAdjacencyTreeNode) serialize() []LinkedAdjacencyTreeNode {
    serializedNodes := LinkedAdjacencyTreeNodes{Nodes: []LinkedAdjacencyTreeNode{}}

    root.collect(&serializedNodes)

    return serializedNodes.Nodes
}

func (node *LinkedAdjacencyTreeNode) collect(serializedNodes *LinkedAdjacencyTreeNodes) {
    serializedNodes.Nodes = append(serializedNodes.Nodes, *node)

    for childIndex := range node.Children {
        node.Children[childIndex].collect(serializedNodes)
    }
}