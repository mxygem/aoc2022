package main

import (
	"testing"
)

func TestScoreByOutcome(t *testing.T) {
	testCases := []struct {
		name    string
		them    byte
		outcome byte
		score   int
	}{
		{
			"(A, X) rock:lose:scissors",
			them("rock"), outcome("lose"), 3,
		},
		{
			"(A, Y) rock:draw:rock",
			them("rock"), outcome("draw"), 4,
		},
		{
			"(A, z) rock:win:paper",
			them("rock"), outcome("win"), 8,
		},
		{
			"(B, X) paper:lose:rock",
			them("paper"), outcome("lose"), 1,
		},
		{
			"(B, Y) paper:draw:paper",
			them("paper"), outcome("draw"), 5,
		},
		{
			"(B, Z) paper:win:scissors",
			them("paper"), outcome("win"), 9,
		},
		{
			"(C, X) scissors:lose:paper",
			them("scissors"), outcome("lose"), 2,
		},
		{
			"(C, Y) scissors:draw:scissors",
			them("scissors"), outcome("draw"), 6,
		},
		{
			"(C, Z) scissors:win:rock",
			them("scissors"), outcome("win"), 7,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := scoreByOutcome(tc.them, tc.outcome)

			if got != tc.score {
				t.Errorf("unexpected score received. got: %d want: %d", got, tc.score)
			}
		})
	}
}

func them(shape string) byte {
	switch shape {
	case "rock":
		return 65
	case "paper":
		return 66
	case "scissors":
		return 67
	default:
		return 0
	}
}

func outcome(res string) byte {
	switch res {
	case "win":
		return 67
	case "draw":
		return 66
	case "lose":
		return 65
	default:
		return 0
	}
}
