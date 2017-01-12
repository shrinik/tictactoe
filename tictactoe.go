package main

import (
	"fmt"
	"strconv"
)

type Position struct {
	Row int
	Col int
}

var board = [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}

func main() {
	Intro()
	ShowBoard()

	// Continue game until the game is won or drawn
	player := "X"
	for winningMove, moves := false, 1; winningMove == false; {
		GetInput(player)
		ShowBoard()

		winningMove = ScoreGame(player)
		if winningMove {
			fmt.Printf("%s is the Winner!! Congratulations!", player)
			break
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}

		moves++
		if moves == 10 {
			fmt.Println("Game was a draw")
			break
		}
	}
}

func Intro() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("Welcome to Tic Tac Toe using Golang!")
	fmt.Println("Developed by Shrinivas Kudva")
	fmt.Println("To play, enter the row and column number of the grid position")
	fmt.Println("---------------------------------------------------------------")
}

func GetInput(player string) {
	var inputStr string
	position := Position{0, 0}
	fmt.Printf("Player %s to play\n", player)

	for isValidInput, errmsg := false, ""; isValidInput == false; {
		isValidInput = false
		fmt.Printf("Enter row number: ")
		fmt.Scanf("%s\n", &inputStr)
		isValidInput, errmsg, position.Row = ValidateInput(inputStr)
		if !isValidInput {
			fmt.Println(errmsg + " Please try again.")
			continue
		}

		fmt.Printf("Enter column number: ")
		fmt.Scanf("%s\n", &inputStr)
		isValidInput, errmsg, position.Col = ValidateInput(inputStr)
		if !isValidInput {
			fmt.Println(errmsg + " Please try again.")
			continue
		}

		// Adjust position to match array index which starts from zero
		position.Row = position.Row - 1
		position.Col = position.Col - 1

		// Check if this position is already in use
		if board[position.Row][position.Col] == "X" ||
			board[position.Row][position.Col] == "O" {
			fmt.Println("This board position already has a value. Please try again")
			isValidInput = false
		} else {
			board[position.Row][position.Col] = player
			isValidInput = true
		}
	}
}

func ValidateInput(inputStr string) (bool, string, int) {
	intVal, err := strconv.Atoi(inputStr)
	if err != nil {
		return false, err.Error(), -1
	}
	switch {
	case intVal < 1 || intVal > 3:
		return false, "Enterd value should be between 1 and 3.", -1
	default:
		return true, "", intVal
	}
}

func ShowBoard() {
	fmt.Printf("\n%s %s %s\n%s %s %s\n%s %s %s\n\n",
		board[0][0], board[0][1], board[0][2],
		board[1][0], board[1][1], board[1][2],
		board[2][0], board[2][1], board[2][2])
}

func ScoreGame(player string) bool {
	var row1, row2, row3, col1, col2, col3, diag1, diag2 int

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			switch {
			case row == 0 && col == 0 && board[row][col] == player:
				row1++
				col1++
				diag1++
			case row == 0 && col == 1 && board[row][col] == player:
				row1++
				col2++
			case row == 0 && col == 2 && board[row][col] == player:
				row1++
				col3++
				diag2++
			case row == 1 && col == 0 && board[row][col] == player:
				row2++
				col1++
			case row == 1 && col == 1 && board[row][col] == player:
				row2++
				col2++
				diag1++
				diag2++
			case row == 1 && col == 2 && board[row][col] == player:
				row2++
				col3++
			case row == 2 && col == 0 && board[row][col] == player:
				row3++
				col1++
				diag2++
			case row == 2 && col == 1 && board[row][col] == player:
				row3++
				col2++
			case row == 2 && col == 2 && board[row][col] == player:
				row3++
				col3++
				diag1++
			}
		}
	}

	if row1 == 3 || row2 == 3 || row3 == 3 || col1 == 3 || col2 == 3 || col3 == 3 ||
		diag1 == 3 || diag2 == 3 {
		return true
	}
	return false
}
