package main

import (
	"testing"
)

func TestCalcScore(t *testing.T) {
	testCases := []struct {
		name     string
		them, us byte
		want     int
	}{
		{
			name: "draw: rock vs rock (A,X)",
			them: them("rock"), us: us("rock"),
			want: 4,
		},
		{
			name: "win: rock vs paper (A,Y)",
			them: them("rock"), us: us("paper"),
			want: 8,
		},
		{
			name: "lose: rock vs scissors (A,Z)",
			them: them("rock"), us: us("scissors"),
			want: 3,
		},
		{
			name: "lose: paper vs rock (B,X)",
			them: them("paper"), us: us("rock"),
			want: 1,
		},
		{
			name: "draw: paper vs paper (B,Y)",
			them: them("paper"), us: us("paper"),
			want: 5,
		},
		{
			name: "win: paper vs scissors (B,Z)",
			them: them("paper"), us: us("scissors"),
			want: 9,
		},
		{
			name: "win: scissors vs rock (C,X)",
			them: them("scissors"), us: us("rock"),
			want: 7,
		},
		{
			name: "lose: scissors vs paper (C,Y)",
			them: them("scissors"), us: us("paper"),
			want: 2,
		},
		{
			name: "draw: scissors vs scissors (C,Z)",
			them: them("scissors"), us: us("scissors"),
			want: 6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calcScore(tc.them, tc.us)

			if got != tc.want {
				t.Errorf("unexpected value received. got: %d want: %d", got, tc.want)
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

func us(shape string) byte {
	return them(shape)
}
