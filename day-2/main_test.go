package main

import (
    "testing"
)

func TestValidateGame(t *testing.T) {
    var tests = []struct{
        input string
        output int
    }{
        {"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
        {"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
        {"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 0},
        {"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 0},
        {"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
        {"Game 47: 18 green, 1 red, 7 blue; 6 blue, 19 green, 1 red; 5 blue, 7 green, 1 red; 1 red, 5 blue, 16 green; 15 green, 3 blue", 0},
    }

    for _, test := range tests {
        test := test
        t.Run(test.input, func(t *testing.T) {
            if got, expected := ValidateGame(test.input), test.output; got != expected {
                t.Fatalf("got = %d, want = %d", got, expected)
            }
        })
    }
}
