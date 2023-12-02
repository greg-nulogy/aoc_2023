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
        {"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "1"},
        {"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", "0"},
        {`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
          Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
          Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
          Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
          Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, "8"},
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
        {"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "48"},
        {"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", "1560"},
        {`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
          Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
          Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
          Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
          Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, "2286"},
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
