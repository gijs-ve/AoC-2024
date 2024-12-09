package day9

import (
	"aoc/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert := assert.New(t)
	utils := utils.NewUtils()
	example := utils.ReadExampleAsString()

	result := Part1(example)
	assert.Equal(1928, result)
}

func Test_Part2(t *testing.T) {
	assert := assert.New(t)
	utils := utils.NewUtils()
	example := utils.ReadExampleAsString()

	result := Part2(example)
	assert.Equal(2858, result)
}
