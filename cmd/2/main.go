package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

// Diff returns the difference between the CubeSet and the other CubeSet
func (c CubeSet) Diff(other CubeSet) CubeSet {
	return CubeSet{
		Red:   c.Red - other.Red,
		Green: c.Green - other.Green,
		Blue:  c.Blue - other.Blue,
	}
}

// GetPower returns the power of the CubeSet
func (c CubeSet) GetPower() int {
	return c.Red * c.Green * c.Blue
}

// IsSubset checks if the CubeSet is a subset of the other CubeSet
func (c CubeSet) IsSubset(other CubeSet) bool {
	return c.Red <= other.Red && c.Green <= other.Green && c.Blue <= other.Blue
}

// String returns a string representation of the CubeSet
func (c CubeSet) String() string {
	return fmt.Sprintf("red: %d, green: %d, blue: %d", c.Red, c.Green, c.Blue)
}

func NewCubeSetsFromString(str string) ([]CubeSet, error) {
	var sets []CubeSet
	cubeSets := strings.Split(str, ";")

	for _, cubeSetString := range cubeSets {
		cubeSetString = strings.TrimSpace(cubeSetString)

		cubes := strings.Split(cubeSetString, ",")
		for _, cubeString := range cubes {

			cubeString = strings.TrimSpace(cubeString)

			cubeInfo := strings.Split(cubeString, " ")

			if len(cubeInfo) != 2 {
				return nil, fmt.Errorf("invalid cube info: %s", cubeString)
			}

			var set CubeSet

			switch cubeInfo[1] {
			case "red":
				redIteration, err := strconv.Atoi(cubeInfo[0])
				if err != nil {
					return nil, fmt.Errorf("invalid cube format: %s", cubeString)
				}

				set.Red += redIteration
			case "green":
				greenIteration, err := strconv.Atoi(cubeInfo[0])
				if err != nil {
					return nil, fmt.Errorf("invalid cube format: %s", cubeString)
				}

				set.Green += greenIteration
			case "blue":
				blueIteration, err := strconv.Atoi(cubeInfo[0])
				if err != nil {
					return nil, fmt.Errorf("invalid cube format: %s", cubeString)
				}

				set.Blue += blueIteration
			default:
				return nil, fmt.Errorf("invalid cube info: %s", cubeString)
			}

			sets = append(sets, set)
		}
	}

	return sets, nil
}

func main() {
	bytesRead, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(bytesRead)

	lines := strings.Split(fileContent, "\n")

	total := 0
	for _, line := range lines {
		_, sets, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
		}

		var minSet CubeSet

		for _, set := range sets {
			if set.Red > minSet.Red {
				minSet.Red = set.Red
			}

			if set.Green > minSet.Green {
				minSet.Green = set.Green
			}

			if set.Blue > minSet.Blue {
				minSet.Blue = set.Blue
			}
		}

		total += minSet.GetPower()
	}

	fmt.Println(total)
}

func parseLine(line string) (int, []CubeSet, error) {
	var id int

	gameData := strings.Split(line, ":")
	if len(gameData) != 2 {
		return 0, nil, fmt.Errorf("invalid game data: %s", line)
	}

	_, err := fmt.Sscanf(gameData[0], "Game %d", &id)
	if err != nil {
		return 0, nil, err
	}

	cubeSets, err := NewCubeSetsFromString(gameData[1])
	if err != nil {
		return 0, nil, err
	}

	return id, cubeSets, nil
}
