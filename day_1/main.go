package main

import (
	"fmt"
	"strings"
	"strconv"
	"aoc_2023/utils"
)

func Part1(input string) (string, error) {
    sum := 0
    lines := strings.Split(input, "\n")

    for _, line := range lines {
        var digits []int
        for _, char := range strings.Split(line, "") {
            number, err := strconv.Atoi(char)
            if err == nil {
                digits = append(digits, number)
            }
        }
        if len(digits) >= 1 {
            sum += digits[0] * 10
            sum += digits[len(digits)-1]
        }
    }
	return fmt.Sprintf("%d", sum), nil
}

func handleWords(line string, characterIndex int) (int, bool) {
    words := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    lineLength := len(line)
    for word, number := range words {
        wordLength := len(word)
        if characterIndex <= lineLength - wordLength {
            if line[characterIndex:characterIndex + wordLength] == word {
                return number, true
            }
        }
    }

    return 0, false
}

func Part2(input string) (string, error) {
    sum := 0
    lines := strings.Split(input, "\n")

    for _, line := range lines {
        var digits []int
        for characterIndex, character := range strings.Split(line, "") {
            number, err := strconv.Atoi(character)
            if err == nil {
                digits = append(digits, number)
            } else {
                number, found := handleWords(line, characterIndex)
                if found {
                    digits = append(digits, number)
                }
            }
        }
        if len(digits) >= 1 {
            sum += digits[0] * 10
            sum += digits[len(digits)-1]
        }
    }
	return fmt.Sprintf("%d", sum), nil
}

func main() {
    utils.RunPart(Part1, Part2)
}
