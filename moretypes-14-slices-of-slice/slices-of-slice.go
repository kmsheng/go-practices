package main

import (
    "fmt"
    "strings"
)

func hasWon(player string, board [][]string) bool {

    // check horizontal lines
    if (player == board[0][0] && board[0][0] == board[0][1] && board[0][1] == board[0][2]) {
        return true
    }
    if (player == board[1][0] && board[1][0] == board[1][1] && board[1][1] == board[1][2]) {
        return true
    }
    if (player == board[2][0] && board[2][0] == board[2][1] && board[2][1] == board[2][2]) {
        return true
    }

    // check vertical lines
    if (player == board[0][0] && board[0][0] == board[1][0] && board[1][0] == board[2][0]) {
        return true
    }
    if (player == board[0][1] && board[0][1] == board[1][1] && board[1][1] == board[2][1]) {
        return true
    }
    if (player == board[0][2] && board[0][2] == board[1][2] && board[1][2] == board[2][2]) {
        return true
    }

    // check diagonal lines
    if (player == board[0][0] && board[0][0] == board[1][1] && board[1][1] == board[2][2]) {
        return true
    }
    if (player == board[0][2] && board[0][2] == board[1][1] && board[1][1] == board[2][0]) {
        return true
    }
    return false
}

func main() {
    // Create a tic-tac-toe board.
    board := [][]string{
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
    }

    // The players take turns.
    board[0][0] = "X"
    board[2][2] = "O"
    board[1][2] = "X"
    board[1][0] = "O"
    board[0][2] = "X"
    board[2][0] = "O"
    board[0][1] = "X"

    for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }

    fmt.Println("X has won? ", hasWon("X", board))
}
