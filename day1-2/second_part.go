package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
    "fmt"
)
func main() {
    l1, l2 := readFile() 
    fmt.Println(calcSimilarity(l1, l2))
}

func readFile() (l1, l2 []int) {
    file, err := os.ReadFile("../day1/file.txt")
    if err != nil {
        log.Fatal(err)
    }

    fileReader := bytes.NewReader(file)

    scanner := bufio.NewScanner(fileReader)
    var lineNumberOne[]int
    var lineNumberTwo[]int

    for scanner.Scan(){
        lineSeparated := strings.Split(scanner.Text(), "   ")
        lineOne, _ := strconv.Atoi(lineSeparated[0])
        lineTwo, _ := strconv.Atoi(lineSeparated[1])
        lineNumberOne = append(lineNumberOne, lineOne)
        lineNumberTwo = append(lineNumberTwo, lineTwo)
    }
    return lineNumberOne, lineNumberTwo
}
func calcSimilarity(l1, l2 []int) int {
    repetition := make(map[int]int) 
    var acc int

    for _, v := range l1 {
        val, exists := repetition[v]
        if exists {
            acc += v * val
            continue
        }
        var repetitionNum int
        for _, j := range l2 {
            if j == v {
                repetitionNum++
            }
        }
        repetition[v] = repetitionNum
        acc += v * repetitionNum
    }
    return acc
}
