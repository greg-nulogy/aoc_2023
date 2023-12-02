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
        {"1abc2", "12"},
        {"pqr3stu8vwx", "38"},
        {"1abc2\npqr3stu8vwx", "50"},
    }

    for index, testData := range tests {
        testName := fmt.Sprintf("Test %d", index)
        t.Run(testName, func(t *testing.T) {
            testOutput, _ := Part1(testData.input)
            if testOutput != testData.expected {
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
        {"two1nine", "29"},
        {"eightwothree", "83"},
        {"zoneight234\n7pqrstsixteen", "90"},
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
