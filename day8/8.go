package day8

import (
	"aoc/utils"
	Utils "aoc/utils"
	"fmt"
)

func Part1(input []string) int {
	return getAntiNodeCount(input, false)
}

func Part2(input []string) int {
	return getAntiNodeCount(input, true)
}

func getAntiNodeCount(input []string, anyDistance bool) int {
	utils := utils.NewUtils()
	grid := utils.MakeGrid(input, "")
	antennas := utils.Remove(utils.FindUniqueCharactersInGrid(grid), ".")
	antiNodePositions := []Utils.XY{}
	maxX := len(grid)
	maxY := len(grid[0])

	for _, antenna := range antennas {
		positions, antennaPositions := getAntiNodePostions(grid, anyDistance, antenna)
		antiNodePositions = append(antiNodePositions, positions...)
		if anyDistance {
			antiNodePositions = append(antiNodePositions, antennaPositions...)
		}
	}

	for i := 0; i < len(antiNodePositions); i++ {
		for j := i + 1; j < len(antiNodePositions); j++ {
			if antiNodePositions[i] == antiNodePositions[j] {
				antiNodePositions = append(antiNodePositions[:j], antiNodePositions[j+1:]...)
				continue
			}
			if antiNodePositions[j].X < 0 || antiNodePositions[j].Y < 0 || antiNodePositions[j].Y >= maxY || antiNodePositions[j].X >= maxX {
				antiNodePositions = append(antiNodePositions[:j], antiNodePositions[j+1:]...)
			}
		}
	}
	return len(antiNodePositions)
}

func findAntiNodePositions(antennaA utils.XY, antennaB utils.XY, offset utils.XY, anyDistance bool, maxX int, maxY int) []utils.XY {
	type XY = Utils.XY
	result := []XY{}

	if !anyDistance {
		if antennaB.X+offset.X == 0 && antennaB.Y+offset.Y == 1 {
			fmt.Print("x and y are 0")
		}
		return append(result, XY{X: antennaB.X + offset.X, Y: antennaB.Y + offset.Y})
	}

	if anyDistance {
		x := offset.X
		y := offset.Y
		for antennaA.X+x <= maxX && antennaA.X+x >= 0 && antennaA.Y+y <= maxY && antennaA.Y+y >= 0 {
			antiNodePosition := findAntiNodePositions(antennaA, XY{X: antennaA.X + x, Y: antennaA.Y + y}, offset, false, maxX, maxY)
			result = append(result, antiNodePosition...)
			y += offset.Y
			x += offset.X
		}

	}

	return result
}

func getAntiNodePostions(grid [][]string, anyDistance bool, character string) ([]utils.XY, []utils.XY) {
	antiNodePositions := []utils.XY{}
	maxX := len(grid)
	maxY := len(grid[0])

	utils := utils.NewUtils()
	antennaPositions := utils.FindCharactersInGrid(grid, character)

	for antennaAIndex, antennaA := range antennaPositions {
		for antennaBIndex, antennaB := range antennaPositions {
			if antennaAIndex == antennaBIndex {
				continue
			}
			offset := utils.GetOffset(antennaA, antennaB)
			foundAntiNodePositions := findAntiNodePositions(antennaA, antennaB, offset, anyDistance, maxX, maxY)
			antiNodePositions = append(antiNodePositions, foundAntiNodePositions...)
		}
	}
	return antiNodePositions, antennaPositions
}
