package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Displaywelcome() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("| In order to draw an X or an O, you have to input |")
	fmt.Println("| the number of the square in which you want it to |")
	fmt.Println("| be drawn. The following shows you which number   |")
	fmt.Println("| refers to which square:                          |")
	fmt.Println("|                                                  |")
	fmt.Println("|                      1 | 2 | 3                   |")
	fmt.Println("|                      4 | 5 | 6                   |")
	fmt.Println("|                      7 | 8 | 9                   |")
	fmt.Println("| Enjoy the game!                                  |")
	fmt.Println("----------------------------------------------------")

}

func Displayboard(board [9]string) {
	fmt.Println(board[0], "|", board[1], "|", board[2])
	fmt.Println(board[3], "|", board[4], "|", board[5])
	fmt.Println(board[6], "|", board[7], "|", board[8])
}

var gameOver bool = false
var board = [...]string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}

func renewBoard(board [9]string) [9]string {
	for a := 0; a < 9; a++ {
		board[a] = "-"
	}
	return board
}

func Checkwin(board [9]string) {

	if (board[0] == board[1] && board[1] == board[2]) && (board[0] != "-") {
		fmt.Println(board[0], "won!")
		gameOver = true

	} else if (board[3] == board[4] && board[4] == board[5]) && (board[3] != "-") {
		fmt.Println(board[3], "won!")
		gameOver = true

	} else if (board[6] == board[7] && board[7] == board[8]) && board[6] != "-" {
		fmt.Println(board[6], "won!")
		gameOver = true

	} else if (board[0] == board[3] && board[3] == board[6]) && board[0] != "-" {
		fmt.Println(board[0], "won!")
		gameOver = true

	} else if (board[1] == board[4] && board[4] == board[7]) && board[1] != "-" {
		fmt.Println(board[1], "won!")
		gameOver = true

	} else if (board[2] == board[5] && board[5] == board[8]) && board[2] != "-" {
		fmt.Println(board[2], "won!")
		gameOver = true

	} else if (board[0] == board[4] && board[4] == board[8]) && board[0] != "-" {
		fmt.Println(board[0], "won!")
		gameOver = true

	} else if (board[2] == board[4] && board[4] == board[6]) && board[2] != "-" {
		fmt.Println(board[2], "won!")
		gameOver = true
	}
}

func play() {
	gameOver = false
	board = renewBoard(board)

	rand.Seed(time.Now().UnixNano())

	startTemp := rand.Float64()

	var firstVal string
	var secondVal string

	if startTemp >= 0.5 {
		firstVal = "X"
		secondVal = "O"
	} else {
		firstVal = "O"
		secondVal = "X"
	}

	fmt.Println(firstVal, "starts!")
	fmt.Println(secondVal, "follows!")
	Displayboard(board)

	for a := 0; a < 9; a++ {

		if !gameOver {
			pos := inputPosition(board)

			if a%2 == 0 {
				board[pos] = firstVal
			} else {
				board[pos] = secondVal
			}

			Displayboard(board)

			if a >= 3 {
				Checkwin(board)
			}
		}
	}

	if !gameOver {
		fmt.Println("It's a draw!")
	}
}

func main() {
	fmt.Println("Hello and welcome to a game of Tic Tac Toe!")
	Displaywelcome()
	play()

	for {
		fmt.Println("Fancy a new game? Input 'Y' if so.")
		var yes string
		fmt.Scan(&yes)

		if yes == "Y" {
			play()
		} else {
			fmt.Println("Bye bye!")
			break
		}
	}

}

func inputPosition(board [9]string) int {

	var position int
	fmt.Scan(&position)

	if position < 1 || position > 9 {
		fmt.Println("Invalid input! Only allowed inputs are numbers in the range 1-9!")
		return inputPosition(board)
	}

	if board[position-1] == "-" {
		return position - 1
	} else {
		fmt.Println("Position", position, "is already taken. Try again!")
		return inputPosition(board)
	}
}
