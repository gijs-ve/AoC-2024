package utils

import (
	"os"
	"strconv"
	"strings"
)

// utils struct
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
