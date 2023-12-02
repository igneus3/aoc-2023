package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    file, err := os.Open("./input.txt")
    check(err)
    defer file.Close()

    result := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result += CalibrationValue(scanner.Text())
    }
    check(scanner.Err())

   fmt.Println(result) 
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var chars = "123456789"
var numberMap = map[string] int {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

func CalibrationValue(line string) int {
    first, last := 0, 0
    firstIdx := strings.IndexAny(line, chars)
    lastIdx := strings.LastIndexAny(line, chars)

    if firstIdx > -1 {
        first = int(line[firstIdx] - '0')
    }

    if lastIdx > -1 {
        last = int(line[lastIdx] - '0')
    }

    for k, v := range numberMap {
        idx := strings.Index(line, k)
        if idx > -1 && idx < firstIdx {
            firstIdx = idx
            first = v

            if last == 0 {
                last = v
                lastIdx = idx
            }
        }

        idx = strings.LastIndex(line, k)
        if idx > -1 && idx > lastIdx {
            lastIdx = idx
            last = v

            if first == 0 {
                first = v
                firstIdx = idx
            }
        }
    }

    return first * 10 + last
}

