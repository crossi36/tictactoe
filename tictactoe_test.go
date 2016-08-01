package tictactoe

import "testing"

func TestString(t *testing.T) {
	want := "   |   |   \n---+---+---\n   |   |   \n---+---+---\n   |   |   \n"
	game := NewGame()
	s := game.String()
	if s != want {
		t.Errorf("String() returns:\n%s\nwant:\n%s", s, want)
	}

	want = "   | O |   \n---+---+---\n   | X |   \n---+---+---\n   |   |   \n"
	game = NewGame()
	game.Play(1, 1)
	game.Play(1, 0)
	s = game.String()
	if s != want {
		t.Errorf("String() returns:\n%s\nwant:\n%s", s, want)
	}
}

func BenchmarkString(b *testing.B) {
	game := NewGame()
	for i := 0; i < b.N; i++ {
		_ = game.String()
	}
}

func TestNewGame(t *testing.T) {
	game := NewGame()
	if len(game.board) != 3 {
		t.Errorf("Expected field height 3, got: %v", len(game.board))
	}

	for _, fields := range game.board {
		if len(fields) != 3 {
			t.Errorf("Expected field width 3, got: %v", len(fields))
		}
	}
}

func BenchmarkNewGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGame()
	}
}

func TestCurrentPlayer(t *testing.T) {
	game := NewGame()
	currentPlayer := game.CurrentPlayer()
	game.Play(0, 0)
	if currentPlayer == game.CurrentPlayer() {
		t.Error("It was the same players turn twice")
	}
}

func TestWinner(t *testing.T) {
	game := NewGame()
	game.Play(0, 0)
	game.Play(0, 1)
	game.Play(1, 0)
	game.Play(1, 1)
	game.Play(2, 0)
	winner := game.Winner()
	if winner != 0 {
		t.Error("Winner() returns wrong player")
	}

	game = NewGame()
	game.Play(0, 1)
	game.Play(1, 1)
	game.Play(0, 2)
	game.Play(0, 0)
	game.Play(2, 2)
	game.Play(1, 2)
	game.Play(1, 0)
	game.Play(2, 0)
	game.Play(2, 1)
	winner = game.Winner()
	if winner != -1 {
		t.Error("Winner() wrongly returns a player")
	}
}

func TestOver(t *testing.T) {
	game := NewGame()
	if game.Over() {
		t.Error("Over() falsely proclaims a game as finished")
	}

	game.Play(0, 0)
	game.Play(0, 1)
	game.Play(1, 0)
	game.Play(1, 1)
	game.Play(2, 0)
	if !game.Over() {
		t.Error("Over() does not recognize a finished game")
	}
}

func TestPlay(t *testing.T) {
	game := NewGame()
	backup := copySliceOfSliceOfInt(game.board)
	game.Play(0, 0)
	if testEqualitySlicesOfSliceOfInt(backup, game.board) {
		t.Error("Play() does not change the pitch")
	}

	err := game.Play(0, 0)
	if err == nil {
		t.Error("Play() does not recognize already marked fields")
	}

	err = game.Play(-1, 0)
	if err == nil {
		t.Error("Play() does not recognize out of bounds coordinates (negative)")
	}

	err = game.Play(0, 4)
	if err == nil {
		t.Error("Play() does not recognize out of bounds coordinates (positive)")
	}
}

func BenchmarkPlay(b *testing.B) {
	game := NewGame()
	for i := 0; i < b.N; i++ {
		game.Play(0, 0)
	}
}

func TestGameIsWon(t *testing.T) {
	// vertically
	game := NewGame()
	game.Play(0, 0)
	game.Play(0, 1)
	game.Play(1, 0)
	game.Play(1, 1)
	game.Play(2, 0)
	if !game.gameOver {
		t.Error("Play() does not recognize when a game is over")
	}
	if !game.hasWinner {
		t.Error("Play() does not recognize a winning move")
	}
	if game.hasWinner && game.Winner() != 0 {
		t.Error("Play() announces wrong player as winner")
	}

	// horizontally
	game = NewGame()
	game.Play(0, 0)
	game.Play(1, 0)
	game.Play(0, 1)
	game.Play(1, 1)
	game.Play(0, 2)
	if !game.gameOver {
		t.Error("Play() does not recognize when a game is over")
	}
	if !game.hasWinner {
		t.Error("Play() does not recognize a winning move")
	}
	if game.hasWinner && game.Winner() != 0 {
		t.Error("Play() announces wrong player as winner")
	}

	// diagonally left top
	game = NewGame()
	game.Play(0, 0)
	game.Play(1, 0)
	game.Play(1, 1)
	game.Play(1, 2)
	game.Play(2, 2)

	if !game.gameOver {
		t.Error("Play() does not recognize when a game is over")
	}
	if !game.hasWinner {
		t.Error("Play() does not recognize a winning move")
	}
	if game.hasWinner && game.Winner() != 0 {
		t.Error("Play() announces wrong player as winner")
	}

	// diagonally right top
	game = NewGame()
	game.Play(0, 2)
	game.Play(1, 0)
	game.Play(1, 1)
	game.Play(1, 2)
	game.Play(2, 0)

	if !game.gameOver {
		t.Error("Play() does not recognize when a game is over")
	}
	if !game.hasWinner {
		t.Error("Play() does not recognize a winning move")
	}
	if game.hasWinner && game.Winner() != 0 {
		t.Error("Play() announces wrong player as winner")
	}

	// no winner
	game = NewGame()
	game.Play(0, 1)
	game.Play(1, 1)
	game.Play(0, 2)
	game.Play(0, 0)
	game.Play(2, 2)
	game.Play(1, 2)
	game.Play(1, 0)
	game.Play(2, 0)
	game.Play(2, 1)
	if !game.gameOver {
		t.Error("Play() does not recognize when a game is over")
	}
	if game.hasWinner {
		t.Error("Play() falsely recognizes a winning move")
	}
}

func TestGameEnded(t *testing.T) {
	game := NewGame()
	game.Play(0, 0)
	game.Play(0, 1)
	game.Play(1, 0)
	game.Play(1, 1)
	game.Play(2, 0)

	backup := copySliceOfSliceOfInt(game.board)
	game.Play(2, 2)
	if !testEqualitySlicesOfSliceOfInt(backup, game.board) {
		t.Error("Play() ignores GameOver flag")
	}
}

func TestFieldValue(t *testing.T) {
	game := NewGame()
	if game.FieldValue(1, 0) != " " {
		t.Error("FieldValue() returns player token for empty field")
	}

	game.Play(1, 0)
	if v := game.FieldValue(1, 0); v != "X" {
		t.Errorf("FieldValue() returns %s, expected: X", v)
	}
}

/* Helper functions for easier checking/copying of slice of slices */

func testEqualitySlicesOfSliceOfInt(a, b [][]int) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func copySliceOfSliceOfInt(in [][]int) [][]int {
	out := [][]int{}
	for i := range in {
		temp := []int{}
		for j := range in[i] {
			temp = append(temp, in[i][j])
		}
		out = append(out, temp)
	}
	return out
}
