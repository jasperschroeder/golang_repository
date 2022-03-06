package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Displayboard(board [9]string) {
	fmt.Println(board[0], "|", board[1], "|", board[2])
	fmt.Println(board[3], "|", board[4], "|", board[5])
	fmt.Println(board[6], "|", board[7], "|", board[8])
}

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

var gameOver bool = false

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
		fmt.Println(board[7], "won!")
		gameOver = true

	} else if (board[2] == board[5] && board[5] == board[8]) && board[2] != "-" {
		fmt.Println(board[6], "won!")
		gameOver = true

	} else if (board[0] == board[4] && board[4] == board[8]) && board[0] != "-" {
		fmt.Println(board[6], "won!")
		gameOver = true

	} else if (board[2] == board[4] && board[4] == board[6]) && board[2] != "-" {
		fmt.Println(board[6], "won!")
		gameOver = true
	}
}

func main() {
	fmt.Println("Hello and welcome to a game of Tic Tac Toe!")
	Displaywelcome()

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

	board := [...]string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}
	Displayboard(board)

	for a := 0; a < 9; a++ {

		if !gameOver {

			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadRune()

			if err != nil {
				fmt.Println(err)
			}

			pos := int(char - 49) // unicode to integer

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
