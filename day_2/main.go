package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"flag"
	"regexp"
	"math"
)

func LoadInput(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Couldn't read file:", filename)
		return "", err
	}
	return string(content), nil
}

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
    var part int
    flag.IntVar(&part, "part", 1, "part 1 or 2")
    flag.Parse()
    fmt.Println("Running part", part)


	input, err := LoadInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	if part == 1 {
        output, err := Part1(input)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Answer: %v", output)
	} else {
	    output, err := Part2(input)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Answer: %v", output)
	}
}
