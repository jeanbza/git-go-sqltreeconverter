package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
    "io/ioutil"
)

func main() {
    fileText := getFileText("test_data.sql")
    rawAdjacencyNodes := extractNodes(fileText)
    linkedAdjacencyNodes := buildLinkedNodes(rawAdjacencyNodes.Nodes)
    attachLeftsAndRights(linkedAdjacencyNodes)
    outputSql(linkedAdjacencyNodes)

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

func extractNodes(fileText string) RawAdjacencyTreeNodes {
    r := regexp.MustCompile(`(\d+),'\w+ \w+',(\w+)`)
    nodeStrings := r.FindAllStringSubmatch(fileText, -1)

    var adjacencyNodes []RawAdjacencyTreeNode

    for _, nodeString := range nodeStrings {
        adjacencyNodes = append(adjacencyNodes, RawAdjacencyTreeNode{
            Id: nodeString[1],
            ParentId: nodeString[2],
        })
    }

    return RawAdjacencyTreeNodes{Nodes: adjacencyNodes}
}

func buildLinkedNodes(rawAdjacencyNodes []RawAdjacencyTreeNode) (root *LinkedAdjacencyTreeNode) {
    var linkedAdjacencyNodesList []LinkedAdjacencyTreeNode
    var rootNodeId string
    var rootNode *LinkedAdjacencyTreeNode

    // Insert all
    for _, rawNode := range rawAdjacencyNodes {
        linkedAdjacencyNodesList = append(linkedAdjacencyNodesList, LinkedAdjacencyTreeNode{Id: rawNode.Id})
    }

    // Link
    for _, rawNode := range rawAdjacencyNodes {
        if rawNode.ParentId != "null" {
            var childIndex, parentIndex int

            // Find the matching parent and children (based on id = childIndex, parentId = parentIndex) in linkedAdjacencyNodesList
            for index, linkedNode := range linkedAdjacencyNodesList {
                if linkedNode.Id == rawNode.Id {
                    childIndex = index
                }

                if linkedNode.Id == rawNode.ParentId {
                    parentIndex = index
                }
            }

            linkedAdjacencyNodesList[parentIndex].Children = append(linkedAdjacencyNodesList[parentIndex].Children, &linkedAdjacencyNodesList[childIndex])
        } else {
            rootNodeId = rawNode.Id
        }
    }

    // Get root node
    for index, linkedNode := range linkedAdjacencyNodesList {
        if linkedNode.Id == rootNodeId {
            rootNode = &linkedAdjacencyNodesList[index]
        }
    }

    return rootNode
}

func attachLeftsAndRights(root *LinkedAdjacencyTreeNode) {
    attachLeftsAndRightsRecursively(root, 0)
}

func attachLeftsAndRightsRecursively(node *LinkedAdjacencyTreeNode, index int) int {
    index++
    node.Left = index

    for childIndex := range node.Children {
        index = attachLeftsAndRightsRecursively(node.Children[childIndex], index)
    }

    index++
    node.Right = index

    return index
}

func outputSql(root *LinkedAdjacencyTreeNode) {
    var outputSql string
    serializedNodes := serialize(root)

    for _, node := range serializedNodes {
        outputSql += fmt.Sprintf("update foo set left = %d, right = %d where id = %s;\n", node.Left, node.Right, node.Id)
    }

    d1 := []byte(outputSql)
    err := ioutil.WriteFile("output_with_lefts_and_rights.sql", d1, 0644)

    if err != nil {
        fmt.Println(err)
    }
}

func serialize(root *LinkedAdjacencyTreeNode) []LinkedAdjacencyTreeNode {
    var serializedNodes *LinkedAdjacencyTreeNodes

    collect(root, serializedNodes)

    return serializedNodes.Nodes
}

func collect(node *LinkedAdjacencyTreeNode, serializedNodes *LinkedAdjacencyTreeNodes) {
    serializedNodes.Nodes = append(serializedNodes.Nodes, *node)

    for childIndex := range node.Children {
        collect(node.Children[childIndex], serializedNodes)
    }
}