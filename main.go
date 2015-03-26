package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
)

func main() {
    file, err := os.Open("test_data.sql")

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

    fmt.Println(fileText)

    r, _ := regexp.Compile(`(\d+),'\w+ \w+',(\w+)`)

    strings := r.FindAllStringSubmatch(fileText, -1)

    for _, elem := range strings {
        fmt.Println(elem)
    }
}