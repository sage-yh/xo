package main

import (
	"fmt"
	"strings"
)

var board [3][3]string
var currentPlayer string

func main() {
	initializeBoard()
	currentPlayer = "X"
	playGame()
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}
}

func playGame() {
	for {
		displayBoard()
		makeMove()
		if isWinner() {
			displayBoard()
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}
		if isBoardFull() {
			displayBoard()
			fmt.Println("It's a draw!")
			break
		}
		switchPlayer()
	}
}

func displayBoard() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d %s\n", i, strings.Join(board[i][:], " "))
	}
}

func makeMove() {
	var row, col int
	fmt.Printf("Player %s's turn. Enter row (0-2): ", currentPlayer)
	fmt.Scan(&row)
	fmt.Printf("Enter column (0-2): ")
	fmt.Scan(&col)

	if isValidMove(row, col) {
		board[row][col] = currentPlayer
	} else {
		fmt.Println("Invalid move. Try again.")
		makeMove()
	}
}

func isValidMove(row, col int) bool {
	return row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == "-"
}

func isWinner() bool {
	return checkRows() || checkColumns() || checkDiagonals()
}

func checkRows() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
	}
	return false
}

func checkColumns() bool {
	for i := 0; i < 3; i++ {
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}
	return false
}

func checkDiagonals() bool {
	return (board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer) ||
		(board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer)
}

func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}
