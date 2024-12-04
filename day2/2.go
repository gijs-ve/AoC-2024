package day2

import (
	"aoc/utils"
)

func Part1(input []string) int {
	return countSafeReports(input)
}

type Report struct {
	line    int
	numbers []string
}

func Part2(input []string) int {

	utils := utils.NewUtils()
	dampenedReports := []Report{}
	safeReports := 0

	for i := range input {
		report := utils.SplitString(input[i], " ")
		isSafe := reportIsSafe(report)
		if !isSafe {
			for r := range report {
				dampenedReport := []string{}
				for j := range report {
					if j == r {
						continue
					}
					dampenedReport = append(dampenedReport, report[j])
				}
				dampenedReports = append(dampenedReports, Report{line: i, numbers: dampenedReport})
			}
		}
		if isSafe {
			safeReports += 1
		}
	}

	safeDampenedReports := []int{}

	for i := range dampenedReports {
		reportIsAlreadySafe := false
		for j := range safeDampenedReports {
			if safeDampenedReports[j] == dampenedReports[i].line {
				reportIsAlreadySafe = true
			}
		}
		if reportIsAlreadySafe {
			continue
		}
		if reportIsSafe(dampenedReports[i].numbers) {
			safeDampenedReports = append(safeDampenedReports, dampenedReports[i].line)
			safeReports += 1
		}
	}
	return safeReports
}

func countSafeReports(input []string) int {
	utils := utils.NewUtils()
	safeReports := 0
	for i := range input {
		report := utils.SplitString(input[i], " ")
		if reportIsSafe(report) {
			safeReports += 1
		}
	}
	return safeReports
}

func reportIsSafe(report []string) bool {
	utils := utils.NewUtils()
	isSafe := true
	isDecreasing := false
	isIncreasing := false

	for n := range report {
		number := utils.StringToNumber(report[n])
		if n == len(report)-1 {
			break
		}
		nextNumber := utils.StringToNumber(report[n+1])
		difference := utils.Abs(number - nextNumber)
		if difference < 1 || difference > 3 {
			isSafe = false
		}
		if number > nextNumber {
			isDecreasing = true
			if isIncreasing {
				isSafe = false
				break
			}
		}
		if number < nextNumber {
			isIncreasing = true
			if isDecreasing {
				isSafe = false
				break
			}
		}
	}
	return isSafe
}
