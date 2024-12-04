package day4

import (
	"aoc/utils"
)

func Part1(input []string) int {
	letters := [][]string{}
	for i := range input {
		yStruct := utils.NewUtils().SplitString(input[i], "")
		letters = append(letters, yStruct)
	}

	xMasCount := 0
	for x := range letters {
		for y := range letters[x] {
			if (y) > len(letters[x]) {
				break
			}
			if letters[x][y] == "X" {
				if CheckPositionForLetter(x-1, y, letters, "M") && CheckPositionForLetter(x-2, y, letters, "A") && CheckPositionForLetter(x-3, y, letters, "S") {
					xMasCount += 1
				}
				if CheckPositionForLetter(x+1, y, letters, "M") && CheckPositionForLetter(x+2, y, letters, "A") && CheckPositionForLetter(x+3, y, letters, "S") {
					xMasCount += 1
				}
				if CheckPositionForLetter(x, y-1, letters, "M") && CheckPositionForLetter(x, y-2, letters, "A") && CheckPositionForLetter(x, y-3, letters, "S") {
					xMasCount += 1
				}
				if CheckPositionForLetter(x, y+1, letters, "M") && CheckPositionForLetter(x, y+2, letters, "A") && CheckPositionForLetter(x, y+3, letters, "S") {
					xMasCount += 1
				}
				if CheckPositionForLetter(x-1, y-1, letters, "M") && CheckPositionForLetter(x-2, y-2, letters, "A") && CheckPositionForLetter(x-3, y-3, letters, "S") {
					xMasCount += 1
				}
				if CheckPositionForLetter(x+1, y+1, letters, "M") && CheckPositionForLetter(x+2, y+2, letters, "A") && CheckPositionForLetter(x+3, y+3, letters, "S") {
					xMasCount += 1
				}
				if CheckPositionForLetter(x-1, y+1, letters, "M") && CheckPositionForLetter(x-2, y+2, letters, "A") && CheckPositionForLetter(x-3, y+3, letters, "S") {
					xMasCount += 1
				}
				if CheckPositionForLetter(x+1, y-1, letters, "M") && CheckPositionForLetter(x+2, y-2, letters, "A") && CheckPositionForLetter(x+3, y-3, letters, "S") {
					xMasCount += 1
				}
			}
		}
	}
	return xMasCount
}

func Part2(input []string) int {
	letters := [][]string{}
	for i := range input {
		yStruct := utils.NewUtils().SplitString(input[i], "")
		letters = append(letters, yStruct)
	}

	xMasCount := 0
	for x := range letters {
		for y := range letters[x] {
			if (y) > len(letters[x]) {
				break
			}
			if letters[x][y] == "A" {
				topLeftToBottomRight := (LeftTopHasLetter(x, y, letters, "M") && RightBottomHasLetter(x, y, letters, "S")) || (LeftTopHasLetter(x, y, letters, "S") && RightBottomHasLetter(x, y, letters, "M"))
				topRightToBottomLeft := (RightTopHasLetter(x, y, letters, "M") && LeftBottomHasLetter(x, y, letters, "S")) || (RightTopHasLetter(x, y, letters, "S") && LeftBottomHasLetter(x, y, letters, "M"))

				if topLeftToBottomRight && topRightToBottomLeft {
					xMasCount += 1
				}

			}
		}
	}
	return xMasCount
}

func LeftTopHasLetter(x int, y int, letters [][]string, letter string) bool {
	return CheckPositionForLetter(x-1, y-1, letters, letter)
}

func RightTopHasLetter(x int, y int, letters [][]string, letter string) bool {
	return CheckPositionForLetter(x+1, y-1, letters, letter)
}

func LeftBottomHasLetter(x int, y int, letters [][]string, letter string) bool {
	return CheckPositionForLetter(x-1, y+1, letters, letter)
}

func RightBottomHasLetter(x int, y int, letters [][]string, letter string) bool {
	return CheckPositionForLetter(x+1, y+1, letters, letter)
}

func CheckPositionForLetter(x int, y int, letters [][]string, letter string) bool {
	if x < 0 || y < 0 || x >= len(letters) || y >= len(letters[0]) {
		return false
	}
	return letters[x][y] == letter
}
