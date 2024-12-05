package day5

import (
	"aoc/utils"
)

func Part1(input []string) int {
	utils := utils.NewUtils()
	updateRules, updates := utils.SplitByEmptyLine(input)
	updateRulePairs := GetUpdateRulePairs(updateRules)
	return GetUpdateMiddleNumbers(updates, updateRulePairs)
}

func Part2(input []string) int {
	utils := utils.NewUtils()
	updateRules, updates := utils.SplitByEmptyLine(input)
	updateRulePairs := GetUpdateRulePairs(updateRules)

	inCorrectUpdates := []string{}

	for _, rule := range updates {
		if !UpdateIsInRightOrder(rule, updateRulePairs) {
			inCorrectUpdates = append(inCorrectUpdates, rule)
		}
	}

	orderedUpdates := []string{}

	for _, rule := range inCorrectUpdates {
		orderedUpdate := PlaceInRightOrder(rule, updateRulePairs)
		orderedUpdates = append(orderedUpdates, orderedUpdate)
	}

	return GetUpdateMiddleNumbers(orderedUpdates, updateRulePairs)
}

func GetUpdateMiddleNumbers(updates []string, updateRulePairs [][]string) int {
	utils := utils.NewUtils()
	totalSum := 0
	for _, rule := range updates {
		if UpdateIsInRightOrder(rule, updateRulePairs) {
			splittedRule := utils.SplitString(rule, ",")
			middleIndex := len(splittedRule) / 2
			totalSum += utils.StringToInt(splittedRule[middleIndex])
		}
	}
	return totalSum
}

func GetUpdateRulePairs(updateRules []string) [][]string {
	utils := utils.NewUtils()
	updateRulePairs := [][]string{}
	for _, rule := range updateRules {
		updateRulePairs = append(updateRulePairs, utils.SplitString(rule, "|"))
	}
	return updateRulePairs
}

func PlaceInRightOrder(input string, updateRulePairs [][]string) string {
	utils := utils.NewUtils()
	numbersToCheck := utils.SplitComma(input)
	newUpdateOrder := numbersToCheck
	for number := range numbersToCheck {
		for rulePair := range updateRulePairs {
			currentRulePair := updateRulePairs[rulePair]
			if currentRulePair[0] == numbersToCheck[number] {
				secondNumberIsInPair := false
				for i := range numbersToCheck {
					if currentRulePair[1] == numbersToCheck[i] {
						secondNumberIsInPair = true
					}
				}
				if !secondNumberIsInPair {
					continue
				}
				secondNumberIsPrintedLater := false
				for i := number + 1; i < len(numbersToCheck); i++ {
					if currentRulePair[1] == numbersToCheck[i] {
						secondNumberIsPrintedLater = true
					}
				}
				if secondNumberIsPrintedLater {
					continue
				}
				for i := range newUpdateOrder {
					if newUpdateOrder[i] == currentRulePair[1] && (i == 0 || newUpdateOrder[i-1] != currentRulePair[0]) {
						newUpdateOrder = append(newUpdateOrder[:i], newUpdateOrder[i+1:]...)
					}
					if newUpdateOrder[i] == currentRulePair[0] {
						newUpdateOrder = append(newUpdateOrder[:i+1], append([]string{currentRulePair[1]}, newUpdateOrder[i+1:]...)...)
					}
				}

			}
		}
	}
	result := utils.JoinByComma(newUpdateOrder)
	if !UpdateIsInRightOrder(result, updateRulePairs) {
		return PlaceInRightOrder(result, updateRulePairs)
	}
	return result
}

func UpdateIsInRightOrder(input string, updateRulePairs [][]string) bool {
	utils := utils.NewUtils()
	numbersToCheck := utils.SplitComma(input)
	appliedRules := [][]string{}
	for number := range numbersToCheck {
		for rulePair := range updateRulePairs {
			currentRulePair := updateRulePairs[rulePair]
			if currentRulePair[1] == numbersToCheck[number] {
				appliedRules = append(appliedRules, currentRulePair)
			}
		}
	}
	for ruleIndex := range appliedRules {
		rule := appliedRules[ruleIndex]
		for number := range numbersToCheck {
			testedNumber := numbersToCheck[number]
			if rule[1] == testedNumber {
				for i := number + 1; i < len(numbersToCheck); i++ {
					if rule[1] == numbersToCheck[i] {
						continue
					}
					if rule[0] == numbersToCheck[i] {
						return false
					}
				}
			}
		}
	}
	return true
}
