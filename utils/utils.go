package utils

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
)

type Utils struct{}

func NewUtils() Utils {
	return Utils{}
}

func (u Utils) ReadSample() []string {
	return readLinesFromFile("sample.txt")
}

func (u Utils) ReadExample() []string {
	return readLinesFromFile("example.txt")
}

func (u Utils) ReadSampleAsString() string {
	return readAsString("sample.txt")
}

func (u Utils) ReadExampleAsString() string {
	return readAsString("example.txt")
}

func readAsString(path string) string {
	fData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fData)
}

func readLinesFromFile(path string) []string {
	fData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(fData), "\n")
}

func (u Utils) JoinBySpace(input []string) string {
	return strings.Join(input, " ")
}

func (u Utils) JoinByComma(input []string) string {
	return strings.Join(input, ",")
}

func (u Utils) SplitByEmptyLine(input []string) ([]string, []string) {
	var groups [][]string
	var group []string

	for _, line := range input {
		if line == "" {
			groups = append(groups, group)
			group = []string{}
		} else {
			group = append(group, line)
		}
	}
	groups = append(groups, group)
	return groups[0], groups[1]
}

func (u Utils) SplitSpaces(str string) []string {
	return strings.Fields(str)
}

func (u Utils) SplitComma(str string) []string {
	return strings.Split(str, ",")
}

func (u Utils) StringToNumber(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func (u Utils) Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (u Utils) Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (u Utils) Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (u Utils) IsEven(x int) bool {
	return x%2 == 0
}

func (u Utils) IsOdd(x int) bool {
	return x%2 == 1
}

func (u Utils) IntLength(x int) int {
	return len(strconv.Itoa(x))
}

func (u Utils) Contains(slice []string, char string) bool {
	for _, c := range slice {
		if c == char {
			return true
		}
	}
	return false
}

func (u Utils) Remove(slice []string, char string) []string {
	newSlice := []string{}
	for _, c := range slice {
		if c != char {
			newSlice = append(newSlice, c)
		}
	}
	return newSlice
}

func (u Utils) SplitString(s string, sep string) []string {
	return strings.Split(s, sep)
}

func (u Utils) StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (u Utils) StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func (u Utils) IntToString(i int) string {
	return strconv.Itoa(i)
}

func (u Utils) MakeGrid(input []string, sep string) [][]string {
	var xy [][]string
	for _, line := range input {
		xy = append(xy, strings.Split(line, sep))
	}
	return xy
}

func (u Utils) FindUniqueCharactersInGrid(input [][]string) []string {
	characterSlice := []string{}
	for _, line := range input {
		for _, character := range line {
			if !u.Contains(characterSlice, character) {
				characterSlice = append(characterSlice, character)
			}
		}
	}
	return characterSlice
}

// grid related
type XY struct {
	X, Y int
}

type Block struct {
	X         int
	Y         int
	Character string
}

func (u Utils) FindCharactersInGrid(grid [][]string, character string) []XY {
	characterSlice := []XY{}
	for x, line := range grid {
		for y, c := range line {
			if c == character {
				characterSlice = append(characterSlice, XY{X: x, Y: y})
			}
		}
	}
	return characterSlice
}

func (u Utils) GetXYDistance(objectA, objectB XY) XY {
	return XY{X: u.Abs(objectA.X - objectB.X), Y: u.Abs(objectA.Y - objectB.Y)}
}

func (u Utils) GetOffset(objectA, objectB XY) XY {
	distance := u.GetXYDistance(objectA, objectB)
	offsetX := distance.X
	offsetY := distance.Y

	if objectA.X > objectB.X {
		offsetX = -distance.X
	}

	if objectA.Y > objectB.Y {
		offsetY = -distance.Y
	}
	return XY{X: offsetX, Y: offsetY}
}

func (u Utils) CopyGrid(grid [][]string) [][]string {
	var newGrid [][]string
	copier.Copy(&newGrid, &grid)
	return newGrid
}

func (u Utils) DeepCopy(src, dst interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}

func (U Utils) AddBounds(grid [][]string, chararacter string) [][]string {
	newGrid := [][]string{}
	for i := 0; i < len(grid)+2; i++ {
		newGrid = append(newGrid, []string{})
		for j := 0; j < len(grid[0])+2; j++ {
			if i == 0 || i == len(grid)+1 || j == 0 || j == len(grid[0])+1 {
				newGrid[i] = append(newGrid[i], chararacter)
			} else {
				newGrid[i] = append(newGrid[i], grid[i-1][j-1])
			}
		}
	}
	return newGrid
}

func (u Utils) FindPositionInGrid(grid [][]string, character string) (int, int, error) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == character {
				return i, j, nil
			}
		}
	}
	return -1, -1, errors.New("Character not found in grid")
}

func (u Utils) GetNextBlock(grid [][]string, x, y int, direction string) string {
	if direction == "up" {
		return grid[x-1][y]
	}
	if direction == "down" {
		return grid[x+1][y]
	}
	if direction == "left" {
		return grid[x][y-1]
	}
	if direction == "right" {
		return grid[x][y+1]
	}
	return ""
}

func (u Utils) GetAdjacentBlocks(grid [][]string, x, y int) []Block {
	return []Block{
		{X: x - 1, Y: y, Character: grid[x-1][y]},
		{X: x + 1, Y: y, Character: grid[x+1][y]},
		{X: x, Y: y - 1, Character: grid[x][y-1]},
		{X: x, Y: y + 1, Character: grid[x][y+1]},
	}
}

func (u Utils) CountInstancesInGrid(grid [][]string, character string) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == character {
				count++
			}
		}
	}
	return count
}

func (u Utils) Factorial(n int) (result int) {
	if n > 0 {
		result = n * u.Factorial(n-1)
		return result
	}
	return 1
}
