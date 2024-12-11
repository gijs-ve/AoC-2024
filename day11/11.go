package day11

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func Part1(input string) int {
	return countStones(input, 25)
}

type Cache struct {
	blinks  int
	stone   int
	results []int
}

func Part2(input string) int {
	return countStones(input, 75)
}

func countStones(input string, maxBlinkCount int) int {
	utils := utils.NewUtils()
	inputStones := getStones(input)
	stones := make(map[string]int)
	for _, index := range inputStones {
		stones[strconv.Itoa(index)]++
	}

	for blinkCount := 1; blinkCount <= maxBlinkCount; blinkCount++ {
		stonesCopy := make(map[string]int)
		for stone, count := range stones {
			if stone == "0" {
				stonesCopy["1"] += count
			}
			if len(stone)%2 == 0 {
				middle := len(stone) / 2
				left := utils.StringToInt((stone[:middle]))
				right := utils.StringToInt(stone[middle:])
				stonesCopy[utils.IntToString(left)] += count
				stonesCopy[utils.IntToString(right)] += count
			}
			if len(stone)%2 != 0 && stone != "0" {
				num := utils.StringToInt(stone)
				stonesCopy[utils.IntToString(num*2024)] += count
			}
		}
		stones = stonesCopy
	}

	result := 0
	for _, count := range stones {
		result += count
	}
	return result
}

// func getStoneCount(stone []int, originalStone int, blinkCount int, maxBlinkCount int, cache *[]Cache, ignoreCache bool) []int {
// 	utils := utils.NewUtils()
// 	var _stone = stone
// 	var _originalStone = originalStone
// 	var _blinkCount = blinkCount

// 	removedBlinks := 0

// 	for _, cacheItem := range *cache {
// 		if (cacheItem.stone == stone[0] && (len(stone) == 1)) && cacheItem.blinks == maxBlinkCount-blinkCount && cacheItem.blinks+blinkCount <= maxBlinkCount {
// 			_stone = cacheItem.results
// 			_blinkCount = cacheItem.blinks + blinkCount
// 			_originalStone = cacheItem.stone
// 			removedBlinks = maxBlinkCount - _blinkCount
// 		}
// 	}

// 	var result = stone
// 	if _blinkCount != maxBlinkCount {
// 		if len(stone) == 1 {
// 			if stone[0] == 0 {
// 				result = getStoneCount(_stone, _originalStone, _blinkCount+1, maxBlinkCount, cache, ignoreCache)
// 			}
// 			if utils.IsEven(utils.IntLength(_stone[0])) {
// 				left, right := splitStone(_stone[0])
// 				result = append(getStoneCount([]int{left}, _originalStone, _blinkCount+1, maxBlinkCount, cache, ignoreCache), getStoneCount([]int{right}, right, _blinkCount+1, maxBlinkCount, cache, ignoreCache)...)
// 			}

// 			if utils.IsOdd(utils.IntLength(_stone[0])) {
// 				result = getStoneCount([]int{_stone[0] * 2024}, _originalStone, _blinkCount+1, maxBlinkCount, cache, ignoreCache)
// 			}
// 		}
// 		if len(stone) > 1 {
// 			for _, _stone := range stone {
// 				cacheResult(cache, _stone, _blinkCount+removedBlinks, maxBlinkCount, stone)
// 				result = append(result, getStoneCount([]int{_stone}, _originalStone, _blinkCount+1, maxBlinkCount, cache, true)...)
// 			}
// 		}
// 	}
// 	isInCache := isInCache(cache, originalStone, blinkCount, maxBlinkCount)
// 	// cache[len(cache)] = Cache{blinks: blinkCount, stone: stone, result: result}
// 	if _blinkCount > 0 && !isInCache && !ignoreCache {
// 		cacheResult(cache, originalStone, blinkCount, maxBlinkCount, stone)
// 	}
// 	if (isInCache && !ignoreCache) && _blinkCount != 0 {
// 		addToCacheResults(cache, originalStone, blinkCount, maxBlinkCount, result)
// 	}
// 	return result
// }

// func isInCache(cache *[]Cache, stone int, blinkCount int, maxBlinkCount int) bool {
// 	for _, cacheItem := range *cache {
// 		if cacheItem.stone == stone && cacheItem.blinks == maxBlinkCount-blinkCount {
// 			return true
// 		}
// 	}
// 	return false
// }

// func addToCacheResults(cache *[]Cache, stone int, blinkCount int, maxBlinkCount int, result []int) {
// 	for i, cacheItem := range *cache {
// 		if cacheItem.stone == stone && cacheItem.blinks == maxBlinkCount-blinkCount {
// 			(*cache)[i].results = append((*cache)[i].results, result...)
// 		}
// 	}
// }

// func cacheResult(cache *[]Cache, stone int, blinkCount int, maxBlinkCount int, result []int) {
// 	*cache = append(*cache, Cache{blinks: blinkCount, stone: stone, results: result})
// }

func getStones(input string) []int {
	utils := utils.NewUtils()
	stones := utils.SplitSpaces(input)
	stoneNumbers := []int{}

	for _, stone := range stones {
		stoneNumbers = append(stoneNumbers, utils.StringToInt(stone))
	}
	return stoneNumbers
}

// func countStones(input string, maxBlinkCount int) int {
// 	stones := getStones(input)
// 	result := applyRuleToStones(stones, 0, maxBlinkCount)
// 	return len(result)
// }

func makeZeroSlice(maxBlinkCount int) []int {
	zeroSlice := make([]int, 0, maxBlinkCount)
	for i := 0; i < maxBlinkCount; i++ {
		zeroSlice[i] = applyRuleToStone([]int{i}, i, maxBlinkCount)[0]
	}
	return zeroSlice
}

func applyRuleToStones(stones []int, blinkCount int, maxBlinkCount int) []int {
	newStones := []int{}
	for _, stone := range stones {
		newStones = append(newStones, applyRuleToStone([]int{stone}, blinkCount, maxBlinkCount)...)
	}
	return newStones
}

func applyRuleToStone(stone []int, blinkCount int, maxBlinkCount int) []int {
	if blinkCount%25 == 0 {
		fmt.Print("\n BC: ", blinkCount, " Stone: ", stone)
	}
	utils := utils.NewUtils()
	if blinkCount == maxBlinkCount {
		return stone
	}
	if len(stone) == 2 {
		return applyRuleToStones(stone, blinkCount, maxBlinkCount)
	}
	increasedBlinkCount := blinkCount + 1
	singleStone := stone[0]
	if singleStone == 0 {
		return applyRuleToStone([]int{1}, increasedBlinkCount, maxBlinkCount)
	}
	if utils.IsEven(utils.IntLength(singleStone)) {
		left, right := splitStone(singleStone)
		return applyRuleToStones([]int{left, right}, increasedBlinkCount, maxBlinkCount)
	}
	return applyRuleToStone([]int{singleStone * 2024}, increasedBlinkCount, maxBlinkCount)
}

func splitStone(stone int) (int, int) {
	utils := utils.NewUtils()
	stoneString := utils.IntToString(stone)
	return utils.StringToInt(stoneString[:len(stoneString)/2]), utils.StringToInt(stoneString[len(stoneString)/2:])
}
