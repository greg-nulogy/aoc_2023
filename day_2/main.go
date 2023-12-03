package main

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"math"
	"aoc_2023/utils"
)

func Part1(input string) (string, error) {
    sum := 0
    lines := strings.Split(input, "\n")
    gameRegexp := regexp.MustCompile(`Game (\d*): (.*)*`)
    cubeDataRegexp := regexp.MustCompile(`(\d*) (red|blue|green)`)

    maxRed := 12
    maxGreen := 13
    maxBlue := 14


    for _, line := range lines {
        if len(line) == 0 {
            continue
        }
        matches := gameRegexp.FindAllStringSubmatch(line, -1)
        gameId, err := strconv.Atoi(matches[0][1])
        if err != nil {
            return "", err
        }
        gameData := matches[0][2]
        selections := strings.Split(gameData, ";")

        validGame := true

        for _, selection := range selections {
            cubesData := strings.Split(selection, ",")
            for _, cubeData := range cubesData {
                cubeMatches := cubeDataRegexp.FindAllStringSubmatch(cubeData, -1)
                number, err := strconv.Atoi(cubeMatches[0][1])
                if err != nil {
                    return "", err
                }
                color := cubeMatches[0][2]

                switch color {
                    case "red":
                        validGame = validGame && number <= maxRed
                    case "green":
                        validGame = validGame && number <= maxGreen
                    case "blue":
                        validGame = validGame && number <= maxBlue
                }
            }
        }

        if validGame {
            sum += gameId
        }

    }
	return fmt.Sprintf("%d", sum), nil
}

func Part2(input string) (string, error) {
    sum := 0
    lines := strings.Split(input, "\n")
    gameRegexp := regexp.MustCompile(`Game \d*: (.*)*`)
    cubeDataRegexp := regexp.MustCompile(`(\d*) (red|blue|green)`)

    for _, line := range lines {
        if len(line) == 0 {
            continue
        }
        matches := gameRegexp.FindAllStringSubmatch(line, -1)
        gameData := matches[0][1]
        selections := strings.Split(gameData, ";")

        minRed := 0.0
        minGreen := 0.0
        minBlue := 0.0

        for _, selection := range selections {
            cubesData := strings.Split(selection, ",")
            for _, cubeData := range cubesData {
                cubeMatches := cubeDataRegexp.FindAllStringSubmatch(cubeData, -1)
                number, err := strconv.Atoi(cubeMatches[0][1])
                if err != nil {
                    return "", err
                }
                color := cubeMatches[0][2]

                switch color {
                    case "red":
                        minRed = math.Max(minRed, float64(number))
                    case "green":
                        minGreen = math.Max(minGreen, float64(number))
                    case "blue":
                        minBlue = math.Max(minBlue, float64(number))
                }
            }
        }
        gamePower := minRed * minBlue * minGreen
        sum = sum + int(gamePower)

    }
	return fmt.Sprintf("%d", sum), nil
}

func main() {
    utils.RunPart(Part1, Part2)
}
