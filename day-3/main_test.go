package main

import (
	"testing"
)

func TestSumPartNumbers(t *testing.T) {
    input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..` 


    want := 4361
    got := SumPartNumbers(input)

    if got != want {
        t.Fatalf("got = %d, wanted = %d", got, want)
    }
}

