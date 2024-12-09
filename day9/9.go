package day9

import (
	"aoc/utils"
)

type DigitCombination struct {
	id         int
	files      int
	space      int
	movedFiles int
	guest      []GuestFile
}

type GuestFile struct {
	id    int
	files int
}

func Part1(input string) int {
	digitCombinations := formDigitCombinations(input)
	position := 0
	checkSum := 0
	leftScan := 0
	rightScan := len(digitCombinations) - 1

	for leftScan < rightScan+1 {
		leftCheck := digitCombinations[leftScan]
		rightCheck := digitCombinations[rightScan]

		for i := 0; i < leftCheck.files; i++ {
			checkSum, position = addToCheckSum(checkSum, leftCheck.id, position)
			digitCombinations[leftScan].files--
		}

		if leftCheck.space == 0 {
			leftScan++
			continue
		}

		if leftCheck.space != 0 && rightCheck.files != 0 {
			movedFiles := 0
			for digitCombinations[leftScan].space > 0 && movedFiles < digitCombinations[rightScan].files {
				checkSum, position = addToCheckSum(checkSum, rightCheck.id, position)
				movedFiles++
				digitCombinations[leftScan].space--
				digitCombinations[rightScan].files--
			}
		}

		if digitCombinations[rightScan].files == 0 {
			rightScan--
			continue
		}

	}

	return checkSum
}

func Part2(input string) int {
	digitCombinations := formDigitCombinations(input)
	rightScan := len(digitCombinations) - 1
	lowestAvailableSpace := 0

	for rightScan > 0 {
		for leftScan := lowestAvailableSpace; leftScan < rightScan; leftScan++ {
			if digitCombinations[leftScan].space >= digitCombinations[rightScan].files {
				currentCombination := digitCombinations[rightScan]
				digitCombinations[leftScan].guest = append(digitCombinations[leftScan].guest, GuestFile{id: currentCombination.id, files: currentCombination.files})
				digitCombinations[leftScan].space -= currentCombination.files
				digitCombinations[rightScan].files = 0
				digitCombinations[rightScan].movedFiles += currentCombination.files
				if (digitCombinations[lowestAvailableSpace]).space == 0 {
					lowestAvailableSpace += 1
				}
				break
			}
		}
		rightScan--
	}

	checkSum := 0
	position := 0
	for i := 0; i < len(digitCombinations); i++ {
		addedCombination := digitCombinations[i]
		if addedCombination.files == 0 {
			if addedCombination.movedFiles > 0 {
				for i := 0; i < addedCombination.movedFiles; i++ {
					checkSum, position = addToCheckSum(checkSum, 0, position)
				}
			} else {
				checkSum, position = addToCheckSum(checkSum, 0, position)
			}

		}
		for i := 0; i < addedCombination.files; i++ {
			checkSum, position = addToCheckSum(checkSum, addedCombination.id, position)
		}
		for j := 0; j < len(digitCombinations[i].guest); j++ {
			guest := addedCombination.guest[j]
			for i := 0; i < guest.files; i++ {
				checkSum, position = addToCheckSum(checkSum, guest.id, position)
			}
		}

		for i := 0; i < addedCombination.space; i++ {
			checkSum, position = addToCheckSum(checkSum, 0, position)
		}
	}

	return checkSum

}

// Old part 2 for first attempt
// func Part2Old(input string) int {
// 	digitCombinations := formDigitCombinations(input)

// 	position := 0
// 	checkSum := 0
// 	leftScan := 0
// 	rightScan := len(digitCombinations) - 1
// 	storedLeftScan := 0
// 	storedPosition := 0

// 	for leftScan <= len(digitCombinations)-1 {
// 		storedLeftCheck := digitCombinations[storedLeftScan]
// 		leftCheck := digitCombinations[leftScan]
// 		rightCheck := digitCombinations[rightScan]

// 		if leftScan == rightScan {
// 			leftScan = storedLeftScan
// 			position = storedPosition
// 			rightScan--
// 			continue
// 		}

// 		if storedPosition != position && rightScan != 0 {
// 			position += leftCheck.files + leftCheck.movedSpace
// 		}

// 		if storedLeftCheck.movedFiles != 1 && leftCheck.movedFiles != 1 {
// 			for i := 0; i < leftCheck.files; i++ {
// 				checkSum, position = addToCheckSum(checkSum, leftCheck.id, position)
// 				digitCombinations[leftScan].files--
// 			}
// 			digitCombinations[leftScan].movedFiles = 1
// 		}

// 		if rightScan == 0 {
// 			leftScan++
// 			storedLeftScan++
// 			position += leftCheck.removedFiles + leftCheck.space + leftCheck.movedSpace
// 			continue
// 		}

// 		if leftCheck.space >= rightCheck.files {
// 			for i := 0; i < rightCheck.files; i++ {
// 				checkSum, position = addToCheckSum(checkSum, rightCheck.id, position)
// 				digitCombinations[leftScan].space -= 1
// 				digitCombinations[leftScan].movedSpace += 1
// 				digitCombinations[rightScan].removedFiles += 1
// 			}
// 			digitCombinations[rightScan].files = 0
// 			if digitCombinations[storedLeftScan].space > 0 {
// 				if storedLeftScan == leftScan {
// 					storedPosition = position
// 				} else {
// 					position = storedPosition
// 				}
// 				for i := storedLeftScan; i <= leftScan; i++ {
// 					if digitCombinations[i].space > 0 {
// 						storedLeftScan = i
// 						break
// 					}
// 				}
// 			}
// 			if digitCombinations[storedLeftScan].space == 0 {
// 				leftScan++
// 				storedLeftScan = leftScan
// 				storedPosition = position
// 			}
// 			rightScan--
// 			continue
// 		}

// 		if leftScan > rightScan {
// 			storedLeftScan = 0
// 			leftScan = storedLeftScan
// 			continue
// 		}

// 		if leftCheck.space < rightCheck.files {
// 			position += leftCheck.space
// 			leftScan++
// 			continue
// 		}

// 		if leftCheck.space != 0 && rightCheck.files != 0 {
// 			movedFiles := 0
// 			for digitCombinations[leftScan].space > 0 && movedFiles < digitCombinations[rightScan].files {
// 				checkSum, position = addToCheckSum(checkSum, rightCheck.id, position)
// 				movedFiles++
// 				digitCombinations[leftScan].space--
// 				digitCombinations[rightScan].files--
// 			}
// 		}

// 		if digitCombinations[rightScan].files == 0 {
// 			rightScan--
// 			continue
// 		}

// 	}
// 	return checkSum
// }

func addToCheckSum(checkSum int, addValue int, position int) (int, int) {
	return checkSum + (position * addValue), position + 1
}

func formDigitCombinations(input string) []DigitCombination {
	utils := utils.NewUtils()
	digitCombinations := []DigitCombination{}
	for i := 0; i < len(input); i += 2 {
		files := utils.StringToInt(string(input[i]))
		space := 0
		if i+1 < len(input) {
			space = utils.StringToInt(string(input[i+1]))
		}
		digitCombinations = append(digitCombinations, DigitCombination{id: i / 2, files: files, space: space})
	}
	return digitCombinations
}
