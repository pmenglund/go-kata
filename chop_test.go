package chop

import "testing"

func TestChop(t *testing.T) {
	testCases := []struct {
		array []int
		find  int
		index int
	}{
		{array: []int{}, find: 3, index: -1},
		{array: []int{1}, find: 3, index: -1},
		{array: []int{1}, find: 1, index: 0},

		{array: []int{1, 3, 5}, find: 1, index: 0},
		{array: []int{1, 3, 5}, find: 3, index: 1},
		{array: []int{1, 3, 5}, find: 5, index: 2},
		{array: []int{1, 3, 5}, find: 0, index: -1},
		{array: []int{1, 3, 5}, find: 2, index: -1},
		{array: []int{1, 3, 5}, find: 4, index: -1},
		{array: []int{1, 3, 5}, find: 6, index: -1},

		{array: []int{1, 3, 5, 7}, find: 1, index: 0},
		{array: []int{1, 3, 5, 7}, find: 2, index: 1},
		{array: []int{1, 3, 5, 7}, find: 3, index: 2},
		{array: []int{1, 3, 5, 7}, find: 5, index: 3},
		{array: []int{1, 3, 5, 7}, find: 0, index: -1},
		{array: []int{1, 3, 5, 7}, find: 2, index: -1},
		{array: []int{1, 3, 5, 7}, find: 4, index: -1},
		{array: []int{1, 3, 5, 7}, find: 6, index: -1},
		{array: []int{1, 3, 5, 7}, find: 8, index: -1},
	}
	for _, test := range testCases {
		if actual := Chop(test.find, test.array); actual != test.index {
			t.Errorf("expected to find %d at index %d, but was %d", test.find, test.index, actual)
		}
	}
}
