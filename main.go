package main

import (
	"fmt"
)

func main() {
	ticTacToeField := [3][3]rune{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}
	activePlayer := 'X'

	for {
		printField(ticTacToeField)

		handleInput(&ticTacToeField, activePlayer)

		winner, ok := determineWinner(ticTacToeField)
		if ok {
			printField(ticTacToeField)

			switch winner {
			case 'X', 'O':
				fmt.Printf("player %c won\n", winner)
			case ' ':
				fmt.Printf("no winner\n")
			}

			return
		}

		switch activePlayer {
		case 'X':
			activePlayer = 'O'
		case 'O':
			activePlayer = 'X'
		}
	}
}

func printField(field [3][3]rune) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if field[i][j] == ' ' {
				fmt.Printf("%d", i*3+j+1)
			} else {
				fmt.Printf("%c", field[i][j])
			}
			if j != 2 {
				fmt.Print("|")
			} else {
				fmt.Print("\n")
			}
		}

		if i != 3 {
			fmt.Println("------")
		}
	}
}

func handleInput(field *[3][3]rune, player rune) {
	var number int

	for {
		fmt.Printf("player %c (1-9): ", player)

		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			fmt.Println("invalid input")

			continue
		}

		if number < 1 || number > 9 {
			fmt.Println("invalid input, not within 1-9")

			continue
		}

		if field[(number-1)/3][(number-1)%3] != ' ' {
			fmt.Println("invalid input, already taken")

			continue
		}

		break
	}

	field[(number-1)/3][(number-1)%3] = player
}

func determineWinner(field [3][3]rune) (rune, bool) {
	for i := 0; i < 3; i++ {
		if field[i][0] != ' ' && field[i][0] == field[i][1] && field[i][0] == field[i][2] {
			return field[i][0], true
		}

		if field[0][i] != ' ' && field[0][i] == field[1][i] && field[0][i] == field[2][i] {
			return field[0][i], true
		}
	}

	if field[0][0] != ' ' && field[0][0] == field[1][1] && field[0][0] == field[2][2] {
		return field[0][0], true
	}

	if field[0][2] != ' ' && field[0][2] == field[1][1] && field[0][2] == field[2][0] {
		return field[0][0], true
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if field[i][j] == ' ' {
				return ' ', false
			}
		}
	}

	return ' ', true
}
