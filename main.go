package main

import (
    "os"
    "fmt"
    "flag"
    "bufio"
    "regexp"
    "io/ioutil"
)

var inputFile = flag.String("input", "", "Specify a .sql input file")
var outputFile = flag.String("output", "", "Specify a .sql output file")

var defaultTarget = "madeUpDb.madeUpTable" 
var target = flag.String("target", defaultTarget, "Specify the table to be altered")

var defaultRegex = `(\w+),(\w+)`
var regex = flag.String("regex", defaultRegex, "Specify the regex to parse the input file with. REQUIRED: the first and second subgroups must be id and parent_id, respectively")

func main() {
    flag.Parse()

    if *inputFile == "" {
        fmt.Println("Please provide an input file with the --input option")
        os.Exit(1)
    }

    if *outputFile == "" {
        fmt.Println("Please provide an output file with the --output option")
        os.Exit(1)
    }

    if *target == defaultTarget {
        fmt.Printf("No target database / table selected. Using %v. Specify with --target\n", *target)
    }

    if *regex == defaultRegex {
        fmt.Printf("No regex specified. Using default %v\n", *regex)
    }

    run(*inputFile, *outputFile)
}

func run(inputFile, outputFile string ) {
    fileText := getFileText(inputFile)
    rawAdjacencyNodes := extractNodes(fileText)
    linkedAdjacencyNodes := buildLinkedNodes(rawAdjacencyNodes.Nodes)
    attachLeftsAndRights(linkedAdjacencyNodes)
    outputSql(linkedAdjacencyNodes, outputFile)
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
        fileText += scanner.Text() + "\n"
    }

    return fileText
}

func extractNodes(fileText string) RawAdjacencyTreeNodes {
    r := regexp.MustCompile(*regex)
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

func buildLinkedNodes(rawAdjacencyNodes []RawAdjacencyTreeNode) (roots []LinkedAdjacencyTreeNode) {
    var linkedAdjacencyNodesList []LinkedAdjacencyTreeNode
    var rootNodeIds []string
    var rootNodes []LinkedAdjacencyTreeNode

    // Insert all
    for _, rawNode := range rawAdjacencyNodes {
        linkedAdjacencyNodesList = append(linkedAdjacencyNodesList, LinkedAdjacencyTreeNode{Id: rawNode.Id})
    }

    // Link
    for _, rawNode := range rawAdjacencyNodes {
        if !rawNode.isRoot() {
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
            rootNodeIds = append(rootNodeIds, rawNode.Id)
        }
    }

    // Get root node
    for index, linkedNode := range linkedAdjacencyNodesList {
        for _, rootNodeId := range rootNodeIds {
            if linkedNode.Id == rootNodeId {
                rootNodes = append(rootNodes, linkedAdjacencyNodesList[index])
            }
        }
    }

    return rootNodes
}

func outputSql(roots []LinkedAdjacencyTreeNode, outputFile string) {
    var outputSql string

    for index := range roots {
        serializedNodes := roots[index].serialize()

        for _, node := range serializedNodes {
            outputSql += fmt.Sprintf("update %s set left = %d, right = %d where id = %s;\n", *target, node.Left, node.Right, node.Id)
        }
    }

    data := []byte(outputSql)
    err := ioutil.WriteFile(outputFile, data, 0644)

    if err != nil {
        fmt.Println(err)
    }
}
