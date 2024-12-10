package day10

import (
	"aoc/utils"
	"fmt"
)

func Part1(input []string) int {
	return GetTrailHeadSum(input, false)
}

func Part2(input []string) int {
	return GetTrailHeadSum(input, true)
}

func GetTrailHeadSum(input []string, countAllTrails bool) int {
	type TrailBlock struct {
		X         int
		Y         int
		Character string
		Origin    int
	}
	utils := utils.NewUtils()
	grid := utils.MakeGrid(input, "")
	grid = utils.AddBounds(grid, "#")
	remainingBlocks := []TrailBlock{}
	zeroBlocks := utils.FindCharactersInGrid(grid, "0")

	for zeroBlockIndex, zeroBlock := range zeroBlocks {
		remainingBlocks = append(remainingBlocks, TrailBlock{
			X:         zeroBlock.X,
			Y:         zeroBlock.Y,
			Character: "0",
			Origin:    zeroBlockIndex,
		})
	}

	currentIndex := 0
	trailheadScoreSum := 0
	for len(remainingBlocks) > 0 {
		if (currentIndex + 1) > len(remainingBlocks) {
			currentIndex = 0
			continue
		}
		currentBlock := remainingBlocks[currentIndex]
		if !countAllTrails {
			restartLoop := false
			for compareIndex, compareBlock := range remainingBlocks {
				if currentBlock.X == compareBlock.X && compareBlock.Y == currentBlock.Y && compareBlock.Origin == currentBlock.Origin && currentIndex != compareIndex {
					remainingBlocks = append(remainingBlocks[:currentIndex], remainingBlocks[currentIndex+1:]...)
					restartLoop = true
				}
			}
			if restartLoop {
				continue
			}
		}
		currentCharacter := utils.StringToInt((currentBlock.Character))
		remainingBlocks = append(remainingBlocks[:currentIndex], remainingBlocks[currentIndex+1:]...)
		if currentCharacter == 9 {
			fmt.Printf("+1 for block %v (%v, %v)\n", currentBlock.Character, currentBlock.X, currentBlock.Y)
			trailheadScoreSum += 1
			continue
		}
		adjecentBlocks := utils.GetAdjacentBlocks(grid, currentBlock.X, currentBlock.Y)
		for _, adjecentBlock := range adjecentBlocks {
			if adjecentBlock.Character == "#" {
				continue
			}

			if utils.StringToInt((adjecentBlock.Character))-1 == utils.StringToInt(currentBlock.Character) {
				fmt.Printf("currentBlock: %v (%v, %v), adjecentBlock: %v (%v, %v)\n", currentBlock.Character, currentBlock.X, currentBlock.Y, adjecentBlock.Character, adjecentBlock.X, adjecentBlock.Y)
				remainingBlocks = append(remainingBlocks, TrailBlock{
					X:         adjecentBlock.X,
					Y:         adjecentBlock.Y,
					Character: adjecentBlock.Character,
					Origin:    currentBlock.Origin,
				})
			}
		}
	}
	return trailheadScoreSum
}
