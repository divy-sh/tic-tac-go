package game

import (
	"testing"
)

// Test NewGame function
func TestNewGame(t *testing.T) {
	size := 3
	g := NewGame(size)

	// Check if the board is initialized with the correct size
	if len(g.Board) != size || len(g.Board[0]) != size {
		t.Errorf("Expected board size %d, but got %d", size, len(g.Board))
	}

	// Check if the board is initialized with all zeroes
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if g.Board[i][j] != 0 {
				t.Errorf("Expected 0 at position (%d,%d), but got %d", i, j, g.Board[i][j])
			}
		}
	}

	// Check initial player
	if g.Player != 1 {
		t.Errorf("Expected initial player to be 1, but got %d", g.Player)
	}
}

// Test LegalMoves function
func TestLegalMoves(t *testing.T) {
	g := NewGame(3)

	moves := g.LegalMoves()
	expectedMoves := 9

	if len(moves) != expectedMoves {
		t.Errorf("Expected %d legal moves, but got %d", expectedMoves, len(moves))
	}

	// Simulate a move and check legal moves again
	g, err := g.Move(0, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	moves = g.LegalMoves()
	expectedMoves = 8

	if len(moves) != expectedMoves {
		t.Errorf("Expected %d legal moves, but got %d", expectedMoves, len(moves))
	}
}

// Test Move function
func TestMove(t *testing.T) {
	g := NewGame(3)

	// Valid move
	g, err := g.Move(0, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if g.Board[0][0] != 1 {
		t.Errorf("Expected board[0][0] to be 1, but got %d", g.Board[0][0])
	}

	// Invalid move (out of bounds)
	_, err = g.Move(3, 3)
	if err == nil {
		t.Fatal("Expected error for out of bounds move, but got none")
	}

	// Invalid move (position already taken)
	_, err = g.Move(0, 0)
	if err == nil {
		t.Fatal("Expected error for move on taken position, but got none")
	}
}

// Test IsGameOver function
func TestIsGameOver(t *testing.T) {
	g := NewGame(3)

	if g.IsGameOver() {
		t.Fatal("Expected game to not be over initially")
	}

	// Simulate a winning scenario
	g, _ = g.Move(0, 0)
	g, _ = g.Move(1, 0)
	g, _ = g.Move(0, 1)
	g, _ = g.Move(1, 1)
	g, _ = g.Move(0, 2)

	if !g.IsGameOver() {
		t.Fatal("Expected game to be over after a winning move")
	}
	if g.GetGameStatus() != 1 {
		t.Errorf("Expected winner to be 1, but got %d", g.GetGameStatus())
	}
}

func TestIsGameOverForDrawnGame(t *testing.T) {
	g := NewGame(3)

	if g.IsGameOver() {
		t.Fatal("Expected game to not be over initially")
	}

	g, _ = g.Move(0, 0)
	g, _ = g.Move(1, 1)
	g, _ = g.Move(2, 2)
	g, _ = g.Move(0, 1)
	g, _ = g.Move(2, 1)
	g, _ = g.Move(2, 0)
	g, _ = g.Move(0, 2)
	g, _ = g.Move(1, 2)
	g, _ = g.Move(1, 0)

	if !g.IsGameOver() {
		t.Fatal("Expected game to be over after a drawing move")
	}
	if g.PrintGameStatus() != "It's a draw." {
		t.Errorf("Expected \"It's a draw.\", but got %s", g.PrintGameStatus())
	}
}

// Test PrintGameStatus function
func TestPrintGameStatus(t *testing.T) {
	g := NewGame(3)

	expected := "Player X's turn"
	if g.PrintGameStatus() != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, g.PrintGameStatus())
	}

	// Simulate moves leading to a win
	g, _ = g.Move(0, 0)
	g, _ = g.Move(1, 0)
	g, _ = g.Move(0, 1)
	g, _ = g.Move(1, 1)
	g, _ = g.Move(0, 2)

	expected = "X wins!"
	if g.PrintGameStatus() != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, g.PrintGameStatus())
	}
}

func TestPrintGameStatusForO(t *testing.T) {
	g := NewGame(3)

	expected := "Player O's turn"

	// Simulate moves leading to a win
	g, _ = g.Move(0, 0)
	if g.PrintGameStatus() != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, g.PrintGameStatus())
	}
	g, _ = g.Move(1, 0)
	g, _ = g.Move(0, 1)
	g, _ = g.Move(1, 1)
	g, _ = g.Move(2, 2)
	g, _ = g.Move(1, 2)

	expected = "O wins!"
	if g.PrintGameStatus() != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, g.PrintGameStatus())
	}
}

// Test PrintBoard function
func TestPrintBoard(t *testing.T) {
	g := NewGame(3)

	expected := "   |   |   \n---|---|---\n   |   |   \n---|---|---\n   |   |   \n"
	if g.PrintBoard() != expected {
		t.Errorf("Expected board:\n%s\nBut got:\n%s", expected, g.PrintBoard())
	}

	// Simulate some moves and check the board again
	g, _ = g.Move(0, 0)
	g, _ = g.Move(1, 1)
	expected = " X |   |   \n---|---|---\n   | O |   \n---|---|---\n   |   |   \n"
	if g.PrintBoard() != expected {
		t.Errorf("Expected board:\n%s\nBut got:\n%s", expected, g.PrintBoard())
	}
}

func TestPushMove(t *testing.T) {
	g := NewGame(3)

	// Test a valid move
	move := Move{s1: 0, s2: 0, player: g.Player}
	g, err := g.PushMove(move)
	if err != nil {
		t.Errorf("Unexpected error for valid move: %v", err)
	}
	if g.Board[0][0] != move.player {
		t.Errorf("Expected Board[0][0] to be %d, but got %d", g.Player, g.Board[0][0])
	}

	// Test an invalid move (out of bounds)
	move = Move{s1: 3, s2: 3, player: g.Player}
	_, err = g.PushMove(move)
	if err == nil {
		t.Error("Expected error for out of bounds move, but got nil")
	}

	// Test an invalid move (cell already occupied)
	move = Move{s1: 0, s2: 0, player: g.Player}
	_, err = g.PushMove(move)
	if err == nil {
		t.Error("Expected error for move on already occupied cell, but got nil")
	}

	// Test move after game is over
	g.Winner = 1
	move = Move{s1: 1, s2: 1, player: g.Player}
	_, err = g.PushMove(move)
	if err == nil {
		t.Error("Expected error for move after game over, but got nil")
	}
}
