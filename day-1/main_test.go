package main

import (
    "testing"
)

func TestCalibrationValue(t *testing.T) {
    tests := []struct {
        input string
        output int
    }{
        {"1234", 14},
        {"one23four", 14},
        {"xyzone23fourxyz", 14},
        {"xyzone23fourxyz", 14},
        {"1", 11},
        {"fiveight", 58},
        {"onenine", 19},
        {"twoeight", 28},
        {"threeseven", 37},
        {"foursix", 46},
        {"five", 55},
        {"23krgjlpone", 21},
        {"kfxone67bzb2", 12},
        {"8jjpseven", 87},
        {"236twoknbxlczgd", 22},
        {"sevensrncljm5zmvvrtthreejjd85twonepvj", 71},
        {"1dgzljrtcndqqrqkgxseventhreessnthree", 13},
        {"s2eight6bhshvmsevensix", 26},
        {"5tpbsrf", 55},
        {"two35kxjtnbhxrmdhbgzeight", 28},
        {"khgdlljfjxt6sevenfour35pxone", 61},
        {"qvztdsix2", 62},
        {"6lsgzmjtjrseven8cnbnpgd", 68},
        {"three1sk4hnine", 39},
        {"sixmqhg5tvbvlhtzxgpfqlzone9", 69},
        {"fljgbjmccvpz67one", 61},
        {"eightwothree", 83},
        {"abc", 0},
        {"two1nine", 29},
        {"eightwothree", 83},
        {"abcone2threexyz", 13},
        {"xtwone3four", 24},
        {"4nineeightseven2", 42},
        {"zoneight234", 14},
        {"7pqrstsixteen", 76},
        {"vggvnhqkjseventwo4onetwonftrnd", 72},
        {"fivevdhgg5two1hpnjlvkeighttwo", 52},
    }

    for _, test := range tests {
        test := test
        t.Run(test.input, func(t *testing.T) {
            if got, expected := CalibrationValue(test.input), test.output; got != expected {
                t.Fatalf("got = %d, want = %d", got, expected)
            }
        })
    }
}
