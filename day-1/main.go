package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
    file, err := os.Open("./input.txt")
    check(err)
    defer file.Close()

    result := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result += calibration_value(replace_word_digits(scanner.Text()))
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

var number_map = map[string] string {
    "one": "o1ne",
    "two": "t2wo",
    "three": "t3hree",
    "four": "f4our",
    "five": "f5ive",
    "six": "s6ix",
    "seven": "s7even",
    "eight": "e8ight",
    "nine": "n9ine",
}

func replace_word_digits(input string) string {
    sub_pattern := "one|two|three|four|five|six|seven|eight|nine"
    r_first := regexp.MustCompile(fmt.Sprintf("(%s).*", sub_pattern))
    r_last := regexp.MustCompile(fmt.Sprintf(".*(%s)", sub_pattern))

    result := input
    first_matches := r_first.FindStringSubmatch(result)
    last_matches := r_last.FindStringSubmatch(result)

    if len(first_matches) > 0 {
        first := first_matches[1]
        result = strings.Replace(result, first, number_map[first], 1)
    }

    if len(last_matches) > 0 {
        last := last_matches[1]
        result = strings.ReplaceAll(result, last, number_map[last])
    }

    return result
}
