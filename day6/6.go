package day6

import (
	"aoc/utils"
	"fmt"
)

func Part1(input []string) int {
	utils := utils.NewUtils()
	grid := utils.MakeGrid(input, "")
	grid = utils.AddBounds(grid, "0")
	grid, _ = handleGuardWalk(grid, false)

	return utils.CountInstancesInGrid(grid, "X")
}

func Part2(input []string) int {
	utils := utils.NewUtils()
	grid := utils.MakeGrid(input, "")
	grid = utils.AddBounds(grid, "0")
	loopPositions := 0
	startX, startY, err := utils.FindPositionInGrid(grid, "^")
	if err != nil {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != "." && grid[i][j] != "X" {
				continue
			}
			fmt.Print("\nChecking position: ", i, j)
			fmt.Print("\nRemaining positions:", (len(grid)*len(grid[i]))-(i*len(grid[i])+j))
			grid[startX][startY] = "^"
			grid[i][j] = "8"
			_, isLoop := handleGuardWalk(grid, true)
			grid[i][j] = "."
			if isLoop {
				loopPositions++
			}
		}
	}

	return loopPositions
}

func getNextPosition(nextBlock string, x, y int, direction string) (string, int, int) {
	if nextBlock == "#" || nextBlock == "8" {
		if direction == "up" {
			return "right", x, y + 1
		}
		if direction == "right" {
			return "down", x + 1, y
		}
		if direction == "down" {
			return "left", x, y - 1
		}
		if direction == "left" {
			return "up", x - 1, y
		}
	}
	if nextBlock == "." || nextBlock == "X" {
		if direction == "up" {
			return "up", x - 1, y
		}
		if direction == "left" {
			return "left", x, y - 1
		}
		if direction == "down" {
			return "down", x + 1, y
		}
		if direction == "right" {
			return "right", x, y + 1
		}
	}
	return nextBlock, x, y
}

func handleGuardWalk(grid [][]string, checkForLoops bool) ([][]string, bool) {
	utils := utils.NewUtils()
	guardOutOfbounds := false
	currentDirection := "up"

	knownPositions := []struct {
		x         int
		y         int
		direction string
	}{}

	for !guardOutOfbounds {
		currentX, currentY, err := utils.FindPositionInGrid(grid, "^")
		if err != nil {
			break
		}

		if checkForLoops {
			for _, pos := range knownPositions {
				if pos.x == currentX && pos.y == currentY && pos.direction == currentDirection {
					grid[currentX][currentY] = "X"
					return grid, true
				}
			}
			knownPositions = append(knownPositions, struct {
				x         int
				y         int
				direction string
			}{currentX, currentY, currentDirection})
		}

		newPosition, newX, newY := handleNextPosition(grid, currentX, currentY, currentDirection)
		currentDirection = newPosition
		grid[currentX][currentY] = "X"
		if checkForLoops {
			grid[currentX][currentY] = "."
		}

		grid[newX][newY] = "^"
		if newPosition == "0" {
			grid[newX][newY] = "X"
			guardOutOfbounds = true
			break
		}
	}
	return grid, false
}

func handleNextPosition(grid [][]string, x, y int, direction string) (string, int, int) {
	utils := utils.NewUtils()
	nextBlock := utils.GetNextBlock(grid, x, y, direction)
	newDirection, newX, newY := getNextPosition(nextBlock, x, y, direction)
	if (nextBlock == "#" || nextBlock == "8") && direction != newDirection {
		followingBlock := utils.GetNextBlock(grid, x, y, newDirection)
		if followingBlock == "#" || followingBlock == "8" {
			return handleNextPosition(grid, x, y, newDirection)
		}
	}
	return newDirection, newX, newY
}
