package main

import (
	"fmt"
)

var board = [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}

func main () {
	Intro()
	ShowBoard()
	
	// Continue game until the game is won or drawn
	player := "X"
	for winningMove, moves := false, 1; winningMove == false; {		
	
		ProcessPlayerInput(player)		
		ShowBoard()		
		
		winningMove = ScoreGame(player)
		if winningMove {
			fmt.Printf("%s is the Winner!! Congratulations!", player)
			break
		}	
		
		if player == "X"	{
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
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Welcome to Tic Tac Toe using Golang!")
	fmt.Println("Developed by Shrinivas Kudva")
	fmt.Println("To play, enter the grid position for X or O.")
	fmt.Println("To input a value in the middle of the grid, enter row as 2 and column as 2")
	fmt.Println("---------------------------------------------------------------------------")
}

func ProcessPlayerInput (player string) {
	var row, column int
	
	fmt.Printf("Player %s to play\n", player)
	
	for validentry := false; validentry == false; {
		fmt.Printf("Enter row number: ")	
		fmt.Scanf("%d\n", &row)
		fmt.Printf("Enter column number: ")	
		fmt.Scanf("%d\n", &column)
				
		// Check if this position is already in use
		if (board[row - 1][column - 1] == "X" || board[row - 1][column - 1] == "O") {
			fmt.Println("This position was already used. Please try again")
		} else {
			board[row - 1][column - 1] = player
			validentry = true
		}
	}	
}

func ShowBoard() {	
	fmt.Printf("%s %s %s\n%s %s %s\n%s %s %s\n\n\n", 
		board[0][0], board[0][1], board[0][2], 
		board[1][0], board[1][1], board[1][2], 
		board[2][0], board[2][1], board[2][2])
}

func ScoreGame(player string) bool {
	var row1, row2, row3, col1, col2, col3, diag1, diag2 int
	
	for row:=0; row < 3; row++ {
		for col:=0; col < 3; col++ {
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
	
	if (row1 == 3 || row2 == 3 || row3 == 3 || col1 == 3 || col2 == 3 || col3 == 3 || 
		diag1 == 3 || diag2 == 3) {
		return true
	}
	return false
		
}