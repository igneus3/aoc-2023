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
    cubeCount := map[string] int {
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    
    splitLine := strings.Split(line, ":")
    draws := strings.Split(splitLine[1], ";")

    for _, draw := range draws {
        colors := strings.Split(draw, ",")
        for _, color := range colors {
            splitColor := strings.Split(color, " ")
            count, _ := strconv.Atoi(splitColor[1])
            
            if count > cubeCount[splitColor[2]] {
                cubeCount[splitColor[2]] = count
            }
        }
    }

    return cubeCount["red"] * cubeCount["green"] * cubeCount["blue"]
}
