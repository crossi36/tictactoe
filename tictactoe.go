package tictactoe

import (
	"bytes"
	"fmt"
)

const (
	standardFieldValue = 4
	dimension          = 3
)

// Game represents a game of Tic-Tac-Toe. Obtain it by invoking the factory function
// NewGame().
type Game struct {
	board         [][]int
	currentPlayer int
	moveCount     int
	gameOver      bool
	hasWinner     bool
}

func (g Game) String() string {
	result := bytes.NewBufferString(fmt.Sprintf(" %s | %s | %s\n", mapValue(g.board[0][0]), mapValue(g.board[0][1]), mapValue(g.board[0][2])))
	result.WriteString("---+---+---\n")
	result.WriteString(fmt.Sprintf(" %s | %s | %s\n", mapValue(g.board[1][0]), mapValue(g.board[1][1]), mapValue(g.board[1][2])))
	result.WriteString("---+---+---\n")
	result.WriteString(fmt.Sprintf(" %s | %s | %s\n", mapValue(g.board[2][0]), mapValue(g.board[2][1]), mapValue(g.board[2][2])))
	return result.String()
}

// NewGame is a factory function for a new game of Tic-Tac-Toe.
func NewGame() *Game {
	b := [][]int{
		{standardFieldValue, standardFieldValue, standardFieldValue},
		{standardFieldValue, standardFieldValue, standardFieldValue},
		{standardFieldValue, standardFieldValue, standardFieldValue},
	}
	return &Game{board: b}
}

// CurrentPlayer returns the player id (0 or 1) of the player who has to play next.
func (g Game) CurrentPlayer() int {
	return g.currentPlayer
}

// Winner returns the player id (0 or 1) of the player who has won or -1 if there
// is no winner (yet).
func (g Game) Winner() int {
	if !g.hasWinner {
		return -1
	}
	return g.currentPlayer
}

// Over returns true if the game is finished and false otherwise.
func (g Game) Over() bool {
	return g.gameOver
}

// Play takes x and y coordinates (each between 0 and 2) and marks it for the
// current player. If the coordinates are out of bounds or the field is already
// marked it will return an error.
func (g *Game) Play(x, y int) error {
	switch {
	case g.gameOver:
		return fmt.Errorf("Game is already over")
	case x < 0 || x > 2 || y < 0 || y > 2:
		return fmt.Errorf("Invalid coordinates")
	case g.board[y][x] != standardFieldValue:
		return fmt.Errorf("Field already marked")
	}

	g.board[y][x] = g.currentPlayer
	g.moveCount++

	checkGameStatus(g, x, y)
	if !g.gameOver {
		g.currentPlayer = (g.currentPlayer + 1) % 2
	}
	return nil
}

func checkGameStatus(g *Game, x, y int) {
	switch {
	case g.moveCount < 4:
		return
	case g.moveCount > 8:
		g.gameOver = true
	}

	winCase := dimension * g.currentPlayer
	var row, col, dia, rdia int
	for i := 0; i < dimension; i++ {
		row += g.board[y][i]
		col += g.board[i][x]
		dia += g.board[i][i]
		rdia += g.board[i][dimension-(i+1)]
	}

	if checkFor(winCase, row, col, dia, rdia) {
		g.gameOver = true
		g.hasWinner = true
	}
}
