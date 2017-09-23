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
	if len(game.board) != dimension*dimension {
		t.Errorf("Expected board size %v, got: %v", dimension*dimension, len(game.board))
	}

	for _, field := range game.board {
		if field != Nobody {
			t.Errorf("Expected field with value %v, got: %v", Nobody, field)
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
	if winner != Player1 {
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
	if winner != Nobody {
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
	backup := copySliceOfPlayer(game.board)
	game.Play(0, 0)
	if testEqualitySlicesOfPlayer(backup, game.board) {
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
	if game.hasWinner && game.Winner() != Player1 {
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
	if game.hasWinner && game.Winner() != Player1 {
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
	if game.hasWinner && game.Winner() != Player1 {
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
	if game.hasWinner && game.Winner() != Player1 {
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

	backup := copySliceOfPlayer(game.board)
	game.Play(2, 2)
	if !testEqualitySlicesOfPlayer(backup, game.board) {
		t.Error("Play() ignores GameOver flag")
	}
}

func TestFieldValue(t *testing.T) {
	game := NewGame()
	if game.FieldValue(1, 0) != Nobody {
		t.Error("FieldValue() returns player token for empty field")
	}

	game.Play(1, 0)
	if v := game.FieldValue(1, 0); v != Player1 {
		t.Errorf("FieldValue() returns %#v, expected: Player1", v)
	}
}

/* Helper functions for easier checking/copying of slice of slices */

func testEqualitySlicesOfPlayer(a, b []Player) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func copySliceOfPlayer(in []Player) []Player {
	out := []Player{}
	for _, i := range in {
		out = append(out, i)
	}
	return out
}
