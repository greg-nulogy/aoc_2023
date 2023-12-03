package main

import (
    "fmt"
    "testing"
)

func TestPart1(t *testing.T) {
    var tests = []struct {
        input string
        expected string
    }{
        {
`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "4361"},
    }

    for index, testData := range tests {
        testName := fmt.Sprintf("Test %d", index)
        t.Run(testName, func(t *testing.T) {
            testOutput, err := Part1(testData.input)
            if testOutput != testData.expected {
                t.Errorf("%v", err)
                t.Errorf("got %v, want %v", testOutput, testData.expected)
            }
        })
    }
}

func TestPart2(t *testing.T) {
    var tests = []struct {
        input string
        expected string
    }{
         {
`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "467835"},
    }

    for index, testData := range tests {
        testName := fmt.Sprintf("Test %d", index)
        t.Run(testName, func(t *testing.T) {
            testOutput, _ := Part2(testData.input)
            if testOutput != testData.expected {
                t.Errorf("got %v, want %v", testOutput, testData.expected)
            }
        })
    }
}
