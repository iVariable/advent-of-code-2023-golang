package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytesRead, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(bytesRead)

	lines := strings.Split(fileContent, "\n")

	total := 0
	for _, line := range lines {
		number, err := parseNumber(line)
		if err != nil {
			log.Fatal(err)
		}

		total += number

		fmt.Printf("input line: %s | parsed number: %d | total: %d\n", line, number, total)
	}

	fmt.Println(total)
}

func parseNumber(line string) (int, error) {
	firstDigit := -1
	lastDigit := -1

	length := len(line)

	buffer := ""

	//not utf safe, who cares
	for i := 0; i < length; i++ {
		c, err := strconv.Atoi(string(line[i]))
		if err != nil {
			// not a real digit, try parsing
			buffer += string(line[i])

			c, err = parseDigitFromText(buffer)
			if err != nil {
				continue
			}
		}

		firstDigit = c

		break
	}

	if firstDigit == -1 {
		return 0, fmt.Errorf(`input line contains no digits`)
	}

	buffer = ""

	for i := length - 1; i >= 0; i-- {
		c, err := strconv.Atoi(string([]rune(line)[i]))
		if err != nil {
			// not a real digit, try parsing
			buffer = string(line[i]) + buffer

			c, err = parseDigitFromText(buffer)
			if err != nil {
				continue
			}
		}

		lastDigit = c

		break
	}

	num, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))

	return num, nil
}

func parseDigitFromText(line string) (int, error) {
	switch true {
	case strings.Contains(line, "zero"):
		return 0, nil
	case strings.Contains(line, "one"):
		return 1, nil
	case strings.Contains(line, "two"):
		return 2, nil
	case strings.Contains(line, "three"):
		return 3, nil
	case strings.Contains(line, "four"):
		return 4, nil
	case strings.Contains(line, "five"):
		return 5, nil
	case strings.Contains(line, "six"):
		return 6, nil
	case strings.Contains(line, "seven"):
		return 7, nil
	case strings.Contains(line, "eight"):
		return 8, nil
	case strings.Contains(line, "nine"):
		return 9, nil
	}

	return 0, fmt.Errorf(`input line contains no digits`)
}
