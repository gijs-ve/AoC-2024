package day1

import (
	"aoc/utils"
	"sort"
)

func Part1(input []string) int {
	utils := utils.NewUtils()

	leftSlice := []int{}
	rightSlice := []int{}
	for i := 0; i < len(input); i++ {
		splittedLine := utils.SplitSpaces(input[i])
		leftSlice = append(leftSlice, utils.StringToNumber((splittedLine[0])))
		rightSlice = append(rightSlice, utils.StringToNumber((splittedLine[1])))
	}
	sort.Ints(leftSlice)
	sort.Ints(rightSlice)

	distances := []int{}
	for i := 0; i < len(leftSlice); i++ {
		distances = append(distances, utils.Abs(rightSlice[i]-leftSlice[i]))
	}

	totalDistance := 0
	for i := 0; i < len(distances); i++ {
		totalDistance += distances[i]
	}
	return totalDistance
}

func Part2(input []string) int {
	utils := utils.NewUtils()

	leftSlice := []int{}
	rightSlice := []int{}
	for i := 0; i < len(input); i++ {
		splittedLine := utils.SplitSpaces(input[i])
		leftSlice = append(leftSlice, utils.StringToNumber((splittedLine[0])))
		rightSlice = append(rightSlice, utils.StringToNumber((splittedLine[1])))
	}

	similarityScores := []int{}
	for i := 0; i < len(leftSlice); i++ {
		currentNumber := leftSlice[i]
		for j := 0; j < len(rightSlice); j++ {
			if currentNumber == rightSlice[j] {
				similarityScores = append(similarityScores, currentNumber)
			}
		}
	}

	totalSimilarityScore := 0
	for i := 0; i < len(similarityScores); i++ {
		totalSimilarityScore += similarityScores[i]
	}
	return totalSimilarityScore
}
