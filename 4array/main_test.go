package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	assert.Equal(t, Sum(numbers), 15)
}

func TestSumAll(t *testing.T) {
	assert.Equal(t, []int{15, 24}, SumAll([]int{1, 2, 3, 4, 5}, []int{7, 8, 9}))
	assert.Equal(t, []int{15, 0}, SumAll([]int{1, 2, 3, 4, 5}, []int{}))
	assert.Equal(t, []int{15, 0}, SumAll([]int{1, 2, 3, 4, 5}, []int(nil)))
	assert.Equal(t, []int{}, SumAll())
}

func TestSumAllTails(t *testing.T) {
	assert.Equal(t, []int{14, 17}, SumAllTails([]int{1, 2, 3, 4, 5}, []int{7, 8, 9}))
	assert.Equal(t, []int{14, 0}, SumAllTails([]int{1, 2, 3, 4, 5}, []int{}))
	assert.Equal(t, []int{14, 0}, SumAllTails([]int{1, 2, 3, 4, 5}, []int(nil)))
	assert.Equal(t, []int(nil), SumAllTails())
}
