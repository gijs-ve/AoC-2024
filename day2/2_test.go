package day2

import (
	"aoc/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Part1(t *testing.T) {
	assert := assert.New(t)
	utils := utils.NewUtils()
	example := utils.ReadExample()

	result := Part1(example)
	assert.Equal(2, result)
}

func Test_Part2(t *testing.T) {
	assert := assert.New(t)
	utils := utils.NewUtils()
	example := utils.ReadExample()

	result := Part2(example)
	assert.Equal(4, result)
}
