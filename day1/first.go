package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)
func main() {
    l1, l2 := readList()
    l1 = orderList(l1)
    l2 = orderList(l2)
    
    fmt.Println(calcTotalDistance(l1, l2))
}
func readList() (l1, l2 []int) {
    file, err := os.ReadFile("./file.txt")
    if err != nil {
        log.Fatal(err)
    }
    fileThatIsReader := bytes.NewReader(file)
    scanner := bufio.NewScanner(fileThatIsReader)
    
    var listnumberOne []int
    var listnumberTwo []int
    for scanner.Scan() {
        lineSeparated := strings.Split(scanner.Text(), "   ")
        lineOne, _ := strconv.Atoi(lineSeparated[0])
        lineTwo, _ := strconv.Atoi(lineSeparated[1])

        listnumberOne = append(listnumberOne, lineOne)
        listnumberTwo = append(listnumberTwo, lineTwo)
    }
    return listnumberOne, listnumberTwo
}

func calcTotalDistance(l1, l2 []int) int {
    var sum float64
    for i := 0; i < len(l1); i++ {
        sub := float64(l1[i] - l2[i])
        sum += math.Abs(sub)
    }
    return int(sum)
}

func orderList(l []int) []int {
    for i := 0; i < len(l) - 1; i++ {
        for j := 0; j < len(l)-i-1; j++ {
            if l[j+1] < l[j] {
                l[j], l[j+1] = l[j+1], l[j]
            }
        }
    }
    return l
}

