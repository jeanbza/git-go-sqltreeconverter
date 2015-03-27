package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
)

type RawAdjacencyTreeNode struct {
    Id, ParentId string
}

type LinkedAdjacencyTreeNode struct {
    Id string
    ParentId *LinkedAdjacencyTreeNode
}

func main() {
    fileText := getFileText("test_data.sql")
    rawAdjacencyNodes := extractNodes(fileText)
    linkedAdjacencyNodes := buildLinkedNodes(rawAdjacencyNodes)

    fmt.Println(linkedAdjacencyNodes)
}

func getFileText(filePath string) string {
    file, err := os.Open(filePath)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)
    
    var fileText string

    for scanner.Scan() {
        fileText += scanner.Text()
    }

    return fileText
}

func extractNodes(fileText string) []RawAdjacencyTreeNode {
    r := regexp.MustCompile(`(\d+),'\w+ \w+',(\w+)`)
    nodeStrings := r.FindAllStringSubmatch(fileText, -1)

    var adjacencyNodes []RawAdjacencyTreeNode

    for _, nodeString := range nodeStrings {
        adjacencyNodes = append(adjacencyNodes, RawAdjacencyTreeNode{
            Id: nodeString[1],
            ParentId: nodeString[2],
        })
    }

    return adjacencyNodes
}

func buildLinkedNodes(rawAdjacencyNodes []RawAdjacencyTreeNode) []LinkedAdjacencyTreeNode {
    return []LinkedAdjacencyTreeNode{}
}