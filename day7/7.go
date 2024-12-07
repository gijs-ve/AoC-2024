package day7

import (
	"aoc/utils"
	"fmt"
)

type Line struct {
	result  int
	numbers []int
}

func Part1(input []string) int {
	resultSum := 0
	operators := []string{"+", "*"}
	lines := generateLines(input)
	for _, line := range lines {
		if lineHasValidEquation(line, operators) {
			resultSum += line.result
		}
	}
	return resultSum
}

func Part2(input []string) int {
	resultSum := 0
	operators := []string{"+", "*"}
	failedLines := []Line{}
	lines := generateLines(input)
	for _, line := range lines {
		if lineHasValidEquation(line, operators) {
			resultSum += line.result
			continue
		}
		failedLines = append(failedLines, line)
	}
	operators = []string{"+", "*", "||"}
	for lineIndex, line := range failedLines {
		if (lineIndex+1)%10 == 0 {
			fmt.Print("\n Remaining lines: ", len(failedLines)-lineIndex)
		}
		if lineHasValidEquation(line, operators) {
			resultSum += line.result
		}
	}

	return resultSum
}

func generateLines(input []string) []Line {
	utils := utils.NewUtils()
	lines := []Line{}
	for _, line := range input {
		splitted := utils.SplitString(line, ": ")
		numbers := utils.SplitString(splitted[1], " ")
		result := utils.StringToInt(splitted[0])
		numbersInt := []int{}
		for _, number := range numbers {
			numbersInt = append(numbersInt, utils.StringToInt(number))
		}
		lines = append(lines, Line{result: result, numbers: numbersInt})
	}
	return lines
}

func lineHasValidEquation(line Line, possibleOperators []string) bool {
	operators := generateOperators(len(line.numbers)-1, possibleOperators)
	for _, operator := range operators {
		if solveExpression(line.numbers, operator) == line.result {
			return true
		}
	}
	return false
}

func generateOperators(numbersLength int, possibleOperators []string) [][]string {
	if numbersLength == 0 {
		return [][]string{}
	}
	var operators [][]string
	for _, operator := range possibleOperators {
		if numbersLength == 1 {
			operators = append(operators, []string{operator})
		} else {
			for _, subCombo := range generateOperators(numbersLength-1, possibleOperators) {
				operators = append(operators, append([]string{operator}, subCombo...))
			}
		}
	}
	return operators
}

func solveExpression(numbers []int, operators []string) int {
	utils := utils.NewUtils()
	result := numbers[0]
	for i, operator := range operators {
		if operator == "+" {
			result += numbers[i+1]
		}
		if operator == "*" {
			result *= numbers[i+1]
		}
		if operator == "||" {
			stringNumber := utils.IntToString(result) + utils.IntToString(numbers[i+1])
			result = utils.StringToInt(stringNumber)
		}

	}
	return result
}
