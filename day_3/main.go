package main

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"aoc_2023/utils"
)

type Position struct {
    X int
    Y int
}

func (p *Position) Adjacent(test Position) bool {
    if (p.X == test.X || p.X == test.X - 1 || p.X == test.X + 1) && (p.Y == test.Y + 1 || p.Y == test.Y - 1) {
        return true
    } else if p.Y == test.Y && (p.X == test.X - 1 || p.X == test.X + 1) {
        return true
    }
    return false
}

type NumberBox struct {
    Positions []Position
    Value int
    Summed bool
}

func (b *NumberBox) Adjacent(test Position) bool {
    for _, p := range b.Positions {
        if p.Adjacent(test) {
            return true
        }
    }
    return false
}

func Part1(input string) (string, error) {
    lines := strings.Split(input, "\n")

    numberRegexp := regexp.MustCompile(`\d+`)
    symbolRegexp := regexp.MustCompile(`([^\d\.])`)

    var numbers []NumberBox
    var symbols []Position

    for y, line := range lines {
        symbolMatches := symbolRegexp.FindAllStringIndex(line, -1)
        for _, symbolIndex := range symbolMatches {
            symbols = append(symbols, Position{symbolIndex[0], y})
        }

        numberMatches := numberRegexp.FindAllStringSubmatchIndex(line, -1)
        for _, matchData := range numberMatches {
            startIndex := matchData[0]
            endIndex := matchData[1]

            if startIndex != endIndex {
                number, err := strconv.Atoi(line[startIndex:endIndex])
                if err != nil {
                    fmt.Printf("Couldn't convert number at position %d,%d. Value %s\n", startIndex, y, line[startIndex:endIndex])
                    continue
                }

                numberBox := NumberBox{Value: number, Summed: false}
                for x := startIndex; x < endIndex; x += 1 {
                    numberBox.Positions = append(numberBox.Positions, Position{x,y})
                }

                numbers = append(numbers, numberBox)
            }
        }
    }

    sum := 0
    for _, numberBox := range numbers {
        for _, symbol := range symbols {
            if numberBox.Adjacent(symbol) && !numberBox.Summed {
                sum += numberBox.Value
                numberBox.Summed = true
            }
        }
    }

    return fmt.Sprintf("%d", sum), nil
}

func Part2(input string) (string, error) {
    lines := strings.Split(input, "\n")

    numberRegexp := regexp.MustCompile(`\d+`)
    symbolRegexp := regexp.MustCompile(`(\*)`)

    var numbers []NumberBox
    var symbols []Position

    for y, line := range lines {
        symbolMatches := symbolRegexp.FindAllStringIndex(line, -1)
        for _, symbolIndex := range symbolMatches {
            symbols = append(symbols, Position{symbolIndex[0], y})
        }

        numberMatches := numberRegexp.FindAllStringSubmatchIndex(line, -1)
        for _, matchData := range numberMatches {
            startIndex := matchData[0]
            endIndex := matchData[1]

            number, err := strconv.Atoi(line[startIndex:endIndex])
            if err != nil {
                fmt.Printf("Couldn't convert number at position %d,%d. Value %s\n", startIndex, y, line[startIndex:endIndex])
                continue
            }

            numberBox := NumberBox{Value: number, Summed: false}
            for x := startIndex; x < endIndex; x += 1 {
                numberBox.Positions = append(numberBox.Positions, Position{x,y})
            }

            numbers = append(numbers, numberBox)
        }
    }

    sum := 0
    for _, gearPosition := range symbols {
        adjacentNumberIndices := make([]int, 0, 2)
        for numberBoxIndex, numberBox := range numbers {
            if numberBox.Adjacent(gearPosition) {
                adjacentNumberIndices = append(adjacentNumberIndices, numberBoxIndex)
            }
        }
        if len(adjacentNumberIndices) == 2 {
            sum += (numbers[adjacentNumberIndices[0]].Value * numbers[adjacentNumberIndices[1]].Value)
        }
    }

    return fmt.Sprintf("%d", sum), nil
}

func main() {
    utils.RunPart(Part1, Part2)
}
