package day11

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
	assert.Equal(55312, result)
}

func Test_Part2(t *testing.T) {
	assert := assert.New(t)
	utils := utils.NewUtils()
	example := utils.ReadExampleAsString()

	result := Part2(example)
	assert.Equal(65601038650482, result)
}
