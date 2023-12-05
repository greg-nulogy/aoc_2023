package main

import (
	"fmt"
	"strings"
	"strconv"
	"aoc_2023/utils"
)

func double(value int) int {
    return value << 1
}

func add1(value int) int {
    return value + 1
}

func CardValue(card string, valueFunc func(int) int) (id int, value int, valid bool) {
    prefix, cardData, found := strings.Cut(card, ":")
    if !found {
        return 0, 0, false
    }

    idString, found := strings.CutPrefix(prefix, "Card")
    if !found {
        return 0, 0, false
    }

    id, err := strconv.Atoi(strings.Trim(idString, " "))
    if err != nil {
        fmt.Println("Couldn't find a valid card id:", idString)
        return 0, 0, false
    }

    numbers := strings.Split(cardData, " ")
    hitSeparator := false
    cardValue := 0
    winningNumbers := make(map[int]bool)
    for _, numberString := range numbers {
        if len(numberString) == 0 {
            continue
        } else if numberString == "|" {
            hitSeparator = true
            continue
        }

        number, err := strconv.Atoi(numberString)
        if err != nil {
            fmt.Println("Couldn't convert string to number:", numberString)
            continue
        }

        if hitSeparator {
            if winningNumbers[number] {
                if cardValue == 0 {
                    cardValue = 1
                } else {
                    cardValue = valueFunc(cardValue)
                }
            }
        } else {
            winningNumbers[number] = true
        }
    }

    return id, cardValue, true
}


func Part1(input string) (string, error) {
    lines := strings.Split(input, "\n")

    sum := 0
    for _, line := range lines {
        _, cardValue, valid := CardValue(line, double)

        if valid {
            sum += cardValue
        }
    }

    return fmt.Sprintf("%d", sum), nil
}

func Part2(input string) (string, error) {
    lines := strings.Split(input, "\n")

    additionalCardsWon := make(map[int]int)
    const thisCardValue = 1
    sum := 0

    for lineIndex := len(lines) - 1; lineIndex >= 0; lineIndex -= 1 {
        line := lines[lineIndex]

        cardId, cardValue, valid := CardValue(line, add1)
        if !valid {
            continue
        }

        wonCardsValue := cardValue
        for wonCardIndex := cardId + 1; wonCardIndex <= cardId + cardValue; wonCardIndex += 1 {
            wonCardsValue += additionalCardsWon[wonCardIndex]
        }

        additionalCardsWon[cardId] = wonCardsValue
        sum += wonCardsValue + thisCardValue
    }

    return fmt.Sprintf("%d", sum), nil
}

func main() {
    utils.RunPart(Part1, Part2)
}
