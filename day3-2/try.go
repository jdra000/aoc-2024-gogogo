package main 

import (
    "fmt"
    "os"
    "log"
    "strconv"
    "regexp"
    "bufio"
    "bytes"
    "sort"
)

func main(){
    file, err := os.ReadFile("./file.txt")    
    if err != nil {
        log.Fatal(err)
    }

    fileReader := bytes.NewReader(file)
    scanner := bufio.NewScanner(fileReader)
    var acc int

    expr := `mul\((\d+),\s*(\d+)\)`
    doExpr := `do\(\)`
    dontExpr := `don\'t\(\)`

    mulPattern := regexp.MustCompile(expr)
    doPattern := regexp.MustCompile(doExpr)
    dontPattern := regexp.MustCompile(dontExpr)
    
    do := true
    for scanner.Scan(){
        line := scanner.Bytes()
        hash := make(map[int][2]int)

        mulLoc := mulPattern.FindAllIndex(line, -1)
        doLoc := doPattern.FindAllIndex(line, -1)
        dontLoc := dontPattern.FindAllIndex(line, -1)
        
        // 0 -> multiply
        // 1 -> do
        //-1 -> do not
        for _, l := range mulLoc{
            hash[l[0]] = [2]int{l[1], 0}
        }
        for _, l := range doLoc{
            hash[l[0]] = [2]int{l[1], 1}
        }
        for _, l := range dontLoc{
            hash[l[0]] = [2]int{l[1], -1}
        }
        
        var keys []int
        for k := range hash{
           keys = append(keys, k)
        }
        sort.Ints(keys)

        for _, k := range keys{
            start := k
            end := hash[k][0]
            op := hash[k][1]

            if op == 0{
               if do{
                   //fmt.Printf("%s", line[start:end])
                   acc += calcMul(string(line[start:end])) 
               }else{
                    continue
               }
            } else if op == -1{
                   do = false
            } else {
                do = true
            }
        }
    }
    fmt.Println(acc)
}
func calcMul(m string)int{
    pattern := regexp.MustCompile(`\d+`)
    values := pattern.FindAllString(m, -1)

    v1, err := strconv.Atoi(values[0])
    if err != nil {
        log.Fatal(err)
    }
    v2, err := strconv.Atoi(values[1])
    if err != nil {
        log.Fatal(err)
    }
    return v1 * v2
}
