package main

import (
    "testing"
    "io/ioutil"
    "fmt"
)

/* 
        1-0-12
       /     \
    2-1-9  10-2-11
     /
   3-4-8
   /   \
4-3-5 6-5-7
*/
func TestOneRootAcceptance(t *testing.T) {
    inputFileName := "oneRootAcceptanceTestInput.sql"
    outputFileName := "oneRootAcceptanceTestOutput.sql"

    createInputFile(`INSERT INTO 'foo' VALUES (0,'name 0',null),(1,'name 1',0),(2,'name 2',0),(3,'name 3',4),(4,'name 4',1),(5,'name 5',4)`, inputFileName)

    expectedFileContents := `update foo set left = 1, right = 12 where id = 0;
update foo set left = 2, right = 9 where id = 1;
update foo set left = 3, right = 8 where id = 4;
update foo set left = 4, right = 5 where id = 3;
update foo set left = 6, right = 7 where id = 5;
update foo set left = 10, right = 11 where id = 2;
`

    run(inputFileName, outputFileName)

    actualFileContents := readOutputFile(outputFileName)

    if actualFileContents != expectedFileContents {
        t.Errorf("expected:\n%v\n\ngot:\n%v", expectedFileContents, actualFileContents)
    }
}

func createInputFile(content, fileName string) {
    data := []byte(content)
    err := ioutil.WriteFile(fileName, data, 0644)

    if err != nil {
        fmt.Println(err)
    }
}

func readOutputFile(fileName string) string {
    data, err := ioutil.ReadFile(fileName)
    
    if err != nil {
        fmt.Println(err)
    }

    return string(data)
}