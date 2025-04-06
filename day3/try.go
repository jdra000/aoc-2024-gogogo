package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
    "strconv"
)

func main(){
    file, err := os.ReadFile("./file.txt")    
    if err != nil {
        log.Fatal(err)
    }

    fileReader := bytes.NewReader(file)
    scanner := bufio.NewScanner(fileReader)
    var acc int 
    for scanner.Scan(){
        line := scanner.Text()


        expr := `mul\((\d+),\s*(\d+)\)`
        pattern := regexp.MustCompile(expr)
        response := pattern.FindAllString(line, -1)

        //fmt.Printf("%s", response)
        for _, res := range response {
            pattern := regexp.MustCompile(`\d+`)
            values := pattern.FindAllString(res, -1)

            v1, err := strconv.Atoi(values[0])
            if err != nil {
                log.Fatal(err)
            }
            v2, err := strconv.Atoi(values[1])
            if err != nil {
                log.Fatal(err)
            }

            acc += v1*v2
        }
    }
    fmt.Println(acc)
}

