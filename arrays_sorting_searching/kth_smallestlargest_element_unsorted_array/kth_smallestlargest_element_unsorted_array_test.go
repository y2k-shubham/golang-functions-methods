package kth_smallestlargest_element_unsorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	type Input struct {
		arr []int
		i   int
		j   int
	}
	argsIn := []Input{
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 0},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 1},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 1, j: 5},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 1, j: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 3, j: 4},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 6, j: 6},
	}
	argsOutExpected := [][]int{
		[]int{8, 1, 3, 7, 4, 2, 9},
		[]int{1, 8, 3, 7, 4, 2, 9},
		[]int{9, 1, 3, 7, 4, 2, 8},
		[]int{8, 2, 3, 7, 4, 1, 9},
		[]int{8, 9, 3, 7, 4, 2, 1},
		[]int{8, 1, 3, 4, 7, 2, 9},
		[]int{8, 1, 3, 7, 4, 2, 9},
	}
	for idx, arg := range argsIn {
		argOutExpected := argsOutExpected[idx]

		Swap(arg.arr, arg.i, arg.j)
		argOutComputed := arg.arr

		assert.Equal(t, argOutExpected, argOutComputed, "Output at idx=%d mismatches\narr=%v\ni=%d\tj=%d", idx, arg.arr, arg.i, arg.j)
	}
}

func TestPartition(t *testing.T) {
	type Input struct {
		arr []int
		lo  int
		hi  int
	}
	argsIn := []Input{
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, lo: 0, hi: 0},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, lo: 0, hi: 1},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, lo: 0, hi: 3},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, lo: 2, hi: 3},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, lo: 4, hi: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, lo: 4, hi: 4},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, lo: 5, hi: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, lo: 6, hi: 6},
	}
	type Output struct {
		arr      []int
		pivotInd int
	}
	argsOutExpected := []Output{
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, pivotInd: 0},
		{arr: []int{1, 8, 3, 7, 4, 2, 9}, pivotInd: 0},
		{arr: []int{1, 3, 7, 8, 4, 2, 9}, pivotInd: 2},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, pivotInd: 3},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, pivotInd: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, pivotInd: 4},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, pivotInd: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, pivotInd: 6},
	}
	for idx, arg := range argsIn {
		arrOutExpected := argsOutExpected[idx].arr
		pivotIndOutExpected := argsOutExpected[idx].pivotInd

		pivotIndOutComputed := Partition(arg.arr, arg.lo, arg.hi)
		arrOutComputed := arg.arr

		assert.Equal(t, arrOutExpected, arrOutComputed, "Array at idx=%d mismatches\narr=%v\ni=%d\tj=%d", idx, arg.arr, arg.lo, arg.hi)
		assert.Equal(t, pivotIndOutExpected, pivotIndOutComputed, "Pivot at idx=%d mismatches\narr=%v\ni=%d\tj=%d", idx, arg.arr, arg.lo, arg.hi)
	}
}
