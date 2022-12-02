package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	loc := "testdata/simple-invs.txt"
	expected := []string{"1234", "5678", "", "9012", "3456", "7890", "", "1234"}

	actual, err := readInput(loc)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestElfInventories(t *testing.T) {
	testCases := []struct {
		name     string
		inv      []string
		expected elves
	}{
		{
			name:     "nil inv input",
			inv:      nil,
			expected: nil,
		},
		{
			name: "single elf",
			inv:  []string{"1", "1", "1", "1"},
			expected: elves{
				{origPos: 0, total: 4, inv: []string{"1", "1", "1", "1"}},
			},
		},
		{
			name: "three elves already in order",
			inv:  []string{"1", "", "2", "2", "", "3", "3", "3"},
			expected: elves{
				{origPos: 2, total: 9, inv: []string{"3", "3", "3"}},
				{origPos: 1, total: 4, inv: []string{"2", "2"}},
				{origPos: 0, total: 1, inv: []string{"1"}},
			},
		},
		{
			name: "multiple empty strings in between elves",
			inv:  []string{"1", "", "", "", "2", "2", "", "", "3", "3", "3"},
			expected: elves{
				{origPos: 2, total: 9, inv: []string{"3", "3", "3"}},
				{origPos: 1, total: 4, inv: []string{"2", "2"}},
				{origPos: 0, total: 1, inv: []string{"1"}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := elfInventories(tc.inv)

			if len(tc.expected) != len(actual) {
				t.Errorf("unexpected count of elves received. want: %d got: %d", len(tc.expected), len(actual))
				return
			}
			for i, e := range tc.expected {
				a := actual[i]
				if e.origPos != a.origPos {
					t.Errorf("unexpected origPos found for elf %d. want: %#v got: %#v", i, e.origPos, a.origPos)
				}
				if e.total != a.total {
					t.Errorf("unexpected total found for elf %d. want: %#v got: %#v", i, e.total, a.total)
				}
				if len(e.inv) != len(a.inv) {
					t.Errorf("unexpected count of meals in inventory. want: %d got: %d", len(e.inv), len(a.inv))
					return
				}
				for mi, m := range e.inv {
					ai := a.inv[mi]
					if m != ai {
						t.Errorf("unexpected inv value found for position %d. want: %#v got: %#v", mi, m, ai)
					}
				}
			}
		})
	}
}

func TestCountCalories(t *testing.T) {
	testCases := []struct {
		name     string
		inv      []string
		expected int
	}{
		{
			name:     "empty",
			inv:      nil,
			expected: 0,
		},
		{
			name:     "single meal",
			inv:      []string{"2"},
			expected: 2,
		},
		{
			name:     "multiple meals",
			inv:      []string{"2", "4", "6"},
			expected: 12,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := countCalories(tc.inv)

			if tc.expected != actual {
				t.Errorf("unexpected calorie count. want: %d got: %d", tc.expected, actual)
			}
		})
	}
}
