package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("./input.txt")
    check(err)
    defer file.Close()

    result := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result += ValidateGame(scanner.Text())
    }
    check(scanner.Err())

   fmt.Println(result) 

   count := 0
   for i := 1; i < 101; i++ {
       count += i
   }
   fmt.Println(count)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ValidateGame(line string) int {
    maxCubes := map[string] int {
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    
    splitLine := strings.Split(line, ":")
    draws := strings.Split(splitLine[1], ";")

    valid := true
    for _, draw := range draws {
        colors := strings.Split(draw, ",")
        for _, color := range colors {
            splitColor := strings.Split(color, " ")
            count, _ := strconv.Atoi(splitColor[1])
            
            if count > maxCubes[splitColor[2]] {
                valid = false
            }
        }
    }

    result := 0
    if valid {
        gameId, _ := strconv.Atoi(strings.Split(splitLine[0], " ")[1])
        result = gameId
    }
    
    return result
}
