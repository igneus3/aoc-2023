package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
    file, err := os.Open("./input.txt")
    check(err)
    defer file.Close()

    result := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result += calibration_value(scanner.Text())
    }

    check(scanner.Err())

   fmt.Println(result) 
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func calibration_value(line string) int {
    var a, b int = 0, 0
    var has_a = false

    for _, ch := range line {
        if unicode.IsDigit(ch) {
            if !has_a {
                a = int(ch - '0')
                has_a = true
                continue
            }

            b = int(ch - '0')
        }
    }

    if b == 0 {
        b = a
    }

    return a * 10 + b
}
