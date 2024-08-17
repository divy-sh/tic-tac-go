package engine

import (
	"fmt"
	"testing"

	"github.com/divy-sh/tic-tac-go/game"
)

func TestEval(t *testing.T) {
	game := game.NewGame(3)
	move := Eval(game, false)
	game, _ = game.PushMove(*move)
	if game.Board[1][1] != 1 {
		t.Errorf("Expected %d at 1, 1, but got %d", 1, game.Board[1][1])
	}
}

func TestEvalOnCompletedGame(t *testing.T) {
	game := game.NewGame(3)
	isO := false
	for !game.IsGameOver() {
		move := Eval(game, isO)
		game, _ = game.PushMove(*move)
		isO = !isO
	}
	move := Eval(game, isO)
	if move != nil {
		fmt.Println(move)
		t.Errorf("Expected move to be nil, got not nil")
	}
}
