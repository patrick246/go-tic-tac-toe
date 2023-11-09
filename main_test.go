package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetermineWinner(t *testing.T) {
	testcases := []struct {
		Name     string
		Field    [3][3]rune
		GameOver bool
		Winner   rune
	}{{
		Name: "empty board",
		Field: [3][3]rune{
			{' ', ' ', ' '},
			{' ', ' ', ' '},
			{' ', ' ', ' '},
		},
		GameOver: false,
	}, {
		Name: "no winner",
		Field: [3][3]rune{
			{'X', 'O', 'X'},
			{'O', 'O', 'X'},
			{'X', 'X', 'O'},
		},
		GameOver: true,
		Winner:   ' ',
	}, {
		Name: "horizontal x 1",
		Field: [3][3]rune{
			{'X', 'X', 'X'},
			{'O', 'O', 'X'},
			{'X', 'O', 'O'},
		},
		GameOver: true,
		Winner:   'X',
	}, {
		Name: "horizontal x 2",
		Field: [3][3]rune{
			{'O', 'O', 'X'},
			{'X', 'X', 'X'},
			{'X', 'O', 'O'},
		},
		GameOver: true,
		Winner:   'X',
	}, {
		Name: "horizontal x 3",
		Field: [3][3]rune{
			{'O', 'O', 'X'},
			{'X', 'O', 'O'},
			{'X', 'X', 'X'},
		},
		GameOver: true,
		Winner:   'X',
	}, {
		Name: "horizontal o 1",
		Field: [3][3]rune{
			{'O', 'O', 'O'},
			{'X', 'X', 'O'},
			{'X', 'O', 'O'},
		},
		GameOver: true,
		Winner:   'O',
	}, {
		Name: "horizontal o 2",
		Field: [3][3]rune{
			{'O', 'O', 'X'},
			{'O', 'O', 'O'},
			{'X', 'O', 'O'},
		},
		GameOver: true,
		Winner:   'O',
	}, {
		Name: "horizontal o 3",
		Field: [3][3]rune{
			{'O', 'O', 'X'},
			{'X', 'O', 'O'},
			{'O', 'O', 'O'},
		},
		GameOver: true,
		Winner:   'O',
	}}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			winner, ok := determineWinner(testcase.Field)

			require.Equal(t, testcase.GameOver, ok)

			if ok {
				require.Equal(t, testcase.Winner, winner)
			}
		})
	}
}
