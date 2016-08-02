package tictactoe

import (
	"bytes"
	"fmt"
)

const (
	dimension          = 3
	standardFieldValue = dimension + 1
)

// Player is the type returned by CurrentPlayer() and Winner(). Its value is either
// Player1, Player2 or Tie.
type Player int

// Possible values of the Player type.
const (
	Tie = iota - 1
	Player1
	Player2
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
	var result bytes.Buffer
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			result.WriteString(fmt.Sprintf(" %s ", mapValue(g.board[i][j])))
			if j < dimension-1 {
				result.WriteString("|")
			}
		}
		result.WriteString("\n")
		if i < dimension-1 {
			for j := 0; j < dimension; j++ {
				result.WriteString("---")
				if j < dimension-1 {
					result.WriteString("+")
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}

// NewGame is a factory function for a new game of Tic-Tac-Toe.
func NewGame() *Game {
	b := make([][]int, dimension)
	for i := 0; i < dimension; i++ {
		b[i] = make([]int, dimension)
		for j := 0; j < dimension; j++ {
			b[i][j] = standardFieldValue
		}
	}
	return &Game{board: b}
}

// CurrentPlayer returns Player1 or Player2, depending on who has to play next.
func (g Game) CurrentPlayer() Player {
	return Player(g.currentPlayer)
}

// Winner returns Player 1 or Player 2 depending on who has won or Tie if there
// is no winner (yet).
func (g Game) Winner() Player {
	if !g.hasWinner {
		return Tie
	}
	return Player(g.currentPlayer)
}

// Over returns true if the game is finished and false otherwise.
func (g Game) Over() bool {
	return g.gameOver
}

// Play takes x and y coordinates (each between 0 and dimension-1 (which is 2)) and marks the field for the
// current player. If the coordinates are out of bounds or the field is already
// marked it will return an error.
func (g *Game) Play(x, y int) error {
	switch {
	case g.gameOver:
		return fmt.Errorf("Game is already over")
	case x < 0 || x > dimension-1 || y < 0 || y > dimension-1:
		return fmt.Errorf("Invalid coordinates")
	case g.board[y][x] != standardFieldValue:
		return fmt.Errorf("Field already marked")
	}

	g.board[y][x] = g.currentPlayer
	g.moveCount++

	g.checkStatus(x, y)
	if !g.gameOver {
		g.currentPlayer = (g.currentPlayer + 1) % 2
	}
	return nil
}

func (g *Game) checkStatus(x, y int) {
	switch {
	case g.moveCount < dimension+2:
		return
	case g.moveCount == dimension*dimension:
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

// FieldValue returns the token of the user occupying the field or empty string
// if it is empty.
func (g Game) FieldValue(x, y int) string {
	return mapValue(g.board[y][x])
}

/* Helper functions */

func mapValue(x int) string {
	switch x {
	case 0:
		return "X"
	case 1:
		return "O"
	default:
		return " "
	}
}

func checkFor(value int, items ...int) bool {
	for _, item := range items {
		if value == item {
			return true
		}
	}
	return false
}
