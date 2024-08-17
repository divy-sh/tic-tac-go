package controller

import (
	"testing"
)

func TestNewController(t *testing.T) {
	ctrl := NewController()

	if ctrl.isO {
		t.Errorf("Expected isO to be false initially, got %v", ctrl.isO)
	}

	if len(ctrl.board.Board) != 3 {
		t.Errorf("Expected board size to be 3, got %d", len(ctrl.board.Board))
	}
}

func TestRestart(t *testing.T) {
	ctrl := NewController()

	status := ctrl.Restart()

	if len(status.Board) != 3 {
		t.Errorf("Expected board size to be 3 after restart, got %d", len(status.Board))
	}

	if status.GameStatus != "Player X's turn" {
		t.Errorf("Expected 'Player X's turn', got '%s'", status.GameStatus)
	}
}

func TestMove(t *testing.T) {
	ctrl := NewController()

	// Player X makes a move at (0, 0)
	status, err := ctrl.Move(0, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if status.Board[0][0] != 1 {
		t.Errorf("Expected board[0][0] to be 1 (X), got %d", status.Board[0][0])
	}

	// Check if Player O (engine) makes a move
	foundOMove := false
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if status.Board[i][j] == -1 {
				foundOMove = true
				break
			}
		}
	}
	if !foundOMove {
		t.Error("Expected Player O to make a move, but no O found on the board")
	}
}

func TestMove_OnDrawnGame(t *testing.T) {
	ctrl := NewController()

	// Player X makes a move at (0, 0)
	ctrl.Move(0, 0)
	ctrl.Move(1, 1)
	ctrl.Move(2, 2)
	ctrl.Move(0, 1)
	ctrl.Move(2, 1)
	ctrl.Move(2, 0)
	ctrl.Move(0, 2)
	ctrl.Move(1, 2)
	ctrl.Move(1, 0)
	gameStatus, _ := ctrl.Move(1, 1)
	if gameStatus.GameStatus != "It's a draw." {
		t.Errorf("expected game status to be \"It's a draw.\", got %s", gameStatus.GameStatus)
	}
}

func TestMove_InvalidMove(t *testing.T) {
	ctrl := NewController()

	// Player X makes a move at (0, 0)
	_, err := ctrl.Move(0, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Try to make an invalid move at the same position
	_, err = ctrl.Move(0, 0)
	if err == nil {
		t.Error("Expected an error for invalid move, but got nil")
	}
}

func TestMove_OutOfBoundsMove(t *testing.T) {
	ctrl := NewController()

	_, err := ctrl.Move(29, 29)

	if err == nil {
		t.Error("Expected an error for invalid move, but got nil")
	}
}

func TestMove_OutOfBoundsMoveNegative(t *testing.T) {
	ctrl := NewController()

	_, err := ctrl.Move(-1, -1)

	if err == nil {
		t.Error("Expected an error for invalid move, but got nil")
	}
}

func TestSwitchPlayer(t *testing.T) {
	ctrl := NewController()

	// Switch to Player O
	status := ctrl.SwitchPlayer(true)

	if !ctrl.isO {
		t.Errorf("Expected isO to be true after switching to O, got %v", ctrl.isO)
	}

	if len(status.Board) != 3 {
		t.Errorf("Expected board size to be 3 after switching player, got %d", len(status.Board))
	}

	if status.GameStatus != "Player O's turn" {
		t.Errorf("Expected 'Player O's turn', got '%s'", status.GameStatus)
	}
}

func TestGetGameStatus(t *testing.T) {
	ctrl := NewController()

	// Get initial game status
	status := ctrl.GetGameStatus()

	if status.GameStatus != "Player X's turn" {
		t.Errorf("Expected 'Player X's turn', got '%s'", status.GameStatus)
	}

	// Make a move and check status
	ctrl.Move(0, 0)
	status = ctrl.GetGameStatus()

	if status.GameStatus == "" {
		t.Error("Expected a valid game status after move, got an empty string")
	}
}
