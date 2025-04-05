package main

import(
    "fmt"
    "bytes"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
    "math"
)


func main() {
    file, err := os.ReadFile("./file.txt")
    if err != nil {
        log.Fatal(err)
    }

    fileReader := bytes.NewReader(file)
    scanner := bufio.NewScanner(fileReader)
    var validLines int

    for scanner.Scan(){
        line := strings.Split(scanner.Text(), " ")
        if checkLine(line){
            validLines++
        }
    }
    fmt.Println(validLines)
}
func checkOrder(x,y float64) bool {
   return x < 0 && y < 0 || x > 0 && y >0 
}

func checkLine(line []string) bool {
    var diff float64
    var lastDiff float64
    var absDiff float64

    for i:=0; i<len(line)-1; i++ {

        actualVal, forwardVal, err := toFloat(line[i], line[i+1])
        if err != nil {
            log.Fatal(err)
        }

        diff = *actualVal - *forwardVal
        if i == 0 {
            lastDiff = diff
        }
        if !checkOrder(lastDiff, diff){
            return false
        }
        absDiff = math.Abs(diff)
        lastDiff = diff

        if absDiff > 3 {
            return false
        }
    }
    return true
}
func toFloat(l1, l2 string) (x, y *float64, err error) {
    val, err := strconv.ParseFloat(l1, 64)
    if err != nil{
        return nil, nil, err
    }
    oneForward, err := strconv.ParseFloat(l2, 64)
    if err != nil{
        return nil, nil, err
    }
    return &val, &oneForward, nil
}
