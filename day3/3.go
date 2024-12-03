package day3

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) int {
	muls := FindMuls(input, false)
	return MultiplyMuls(muls)
}

func Part2(input string) int {
	muls := FindMuls(input, true)
	return MultiplyMuls(muls)
}

type Mul struct {
	x int
	y int
}

func FindMuls(input string, enableDoAndDonts bool) []Mul {
	utils := utils.NewUtils()
	stringSlice := utils.SplitString(input, "")
	mulSlice := []Mul{}
	do := true
	for i := range stringSlice {
		currentLetter := stringSlice[i]
		fmt.Print(currentLetter)
		if enableDoAndDonts && stringSlice[i] == "d" && stringSlice[i+1] == "o" {
			if stringSlice[i+2] == "(" && stringSlice[i+3] == ")" {
				{
					do = true
					continue
				}
			}
			if stringSlice[i+2] == "n" && stringSlice[i+3] == "'" && stringSlice[i+4] == "t" && stringSlice[i+5] == "(" && stringSlice[i+6] == ")" {
				{
					do = false
					continue
				}
			}
		}
		if stringSlice[i] != "m" || stringSlice[i+1] != "u" || stringSlice[i+2] != "l" || stringSlice[i+3] != "(" || !CharacterIsNumber(stringSlice[i+4]) || (enableDoAndDonts && !do) {
			continue
		}
		x := ""
		y := ""
		findY := false
		for j := i + 4; j < len(stringSlice); j++ {
			currentCharacter := stringSlice[j]
			if !CharacterIsValid(currentCharacter) {
				break
			}
			if currentCharacter == "," {
				findY = true
				continue
			}
			if currentCharacter == ")" {
				if x != "" && y != "" {
					mulSlice = append(mulSlice, Mul{x: utils.StringToNumber(x), y: utils.StringToNumber(y)})
				}
				break
			}
			if !findY {
				x += currentCharacter
				continue
			}
			if findY {
				y += currentCharacter
				continue
			}
		}
	}
	return mulSlice
}

func CharacterIsValid(currentCharacter string) bool {
	if !CharacterIsNumber(currentCharacter) && !strings.ContainsAny(string(currentCharacter), ",)") {
		return false
	}
	return true
}

func CharacterIsNumber(character string) bool {
	_, err := strconv.Atoi(character)
	return err == nil
}

func MultiplyMuls(mulSlice []Mul) int {
	total := 0
	for i := range mulSlice {
		total += (mulSlice[i].x * mulSlice[i].y)
	}
	return total
}
