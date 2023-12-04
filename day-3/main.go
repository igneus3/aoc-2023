package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
    inputFile, _ := os.Open("./input.txt")
    defer inputFile.Close()

    buffer, _ := io.ReadAll(inputFile)

    fmt.Printf("Result: %d\n", SumPartNumbers(string(buffer)))
}

func SumPartNumbers(input string) int {
    grid := [][]string{}
    lines := strings.Split(input, "\n")
    for _, line := range lines {
        temp := []string{}
        for _, byte := range line {
           temp = append(temp, string(byte))
        }
        grid = append(grid, temp)
    }

    rowCount := len(grid)
    colCount := len(grid[0])
    result := 0

    fmt.Println(grid)

    for r := 0; r < rowCount; r++ {
        for c := 0; c < colCount; c++ {
            if len(grid[r]) == 0 {
                continue
            }

            if IsPossibleGear(grid[r][c]) {
                numbers := []int{}
                for _, direction := range directions {
                    row := r + direction[0]
                    col := c + direction[1]

                    if row < 0 || row >= rowCount || col < 0 || col >= colCount {
                        continue
                    }

                    if IsNumber(grid[row][col]) {
                        start := col
                        for col > 0 && IsNumber(grid[row][col - 1]) {
                            col--
                            start = col
                        }

                        temp := ""
                        cur := start
                        for cur < colCount && IsNumber(grid[row][cur]) {
                            temp += grid[row][cur]
                            grid[row][cur] = "."
                            cur++
                        }

                        partNumber, _ := strconv.Atoi(temp)
                        numbers = append(numbers, partNumber)
                    }
                }

                if len(numbers) == 2 {
                    result += numbers[0] * numbers[1]
                }
            }
        }
    }

    return result
}

var directions = [][]int{
    {-1, 0},
    {-1, -1},
    {0, -1},
    {1, -1},
    {1, 0},
    {1, 1},
    {0, 1},
    {-1, 1},
}

func IsNumber(character string) bool {
    _, err := strconv.Atoi(character)
    return err == nil
}

func IsDot(character string) bool {
    return character == "."
}

func IsPossibleGear(character string) bool {
    return character == "*"
}
