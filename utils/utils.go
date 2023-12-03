package utils

import (
	"fmt"
	"os"
	"flag"
)

func LoadInput(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Couldn't read file:", filename)
		return "", err
	}
	return string(content), nil
}

func DeterminePart() int {
    var part int
    flag.IntVar(&part, "part", 1, "part 1 or 2")
    flag.Parse()
    fmt.Println("Running part", part)

    return part
}

func RunPart(part1 func(string) (string, error), part2 func(string) (string, error)) {
    part := DeterminePart()

    input, err := LoadInput("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    if part == 1 {
        output, err := part1(input)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Answer: %v", output)
    } else {
        output, err := part2(input)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Answer: %v", output)
    }
}
