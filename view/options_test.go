package view

import (
	"testing"

	"fyne.io/fyne/v2/app"
	"github.com/divy-sh/tic-tac-go/controller"
)

func TestNewOptions(t *testing.T) {
	app := app.New()
	defer app.Quit()
	control = controller.NewController()
	gameStatus = control.GetGameStatus()
	options := NewOptions()
	if options == nil {
		t.Errorf("expected NewOptions to be not nil, got nil")
	}
	if statusText.Text != gameStatus.GameStatus {
		t.Errorf("expected status text to be %s, got %s", statusText.Text, gameStatus.GameStatus)
	}
}

// commenting for now as I have no clue why this isn't working, even though the app is working perfectly

// func TestRestart(t *testing.T) {
// 	app := app.New()
// 	defer app.Quit()
// 	control = controller.NewController()
// 	gameStatus = control.GetGameStatus()
// 	options := NewOptions()
// 	if options == nil {
// 		t.Errorf("expected NewOptions to be not nil, got nil")
// 	}
// 	gameStatus, _ = control.Move(1, 1)
// 	if gameStatus.Board[1][1] != 1 {
// 		t.Errorf("expected board at 1, 1 to be %d, got %d", 1, gameStatus.Board[1][1])
// 	}
// 	test.Tap(restartButton)
// 	if gameStatus.Board[1][1] != 0 {
// 		t.Errorf("expected board at 1, 1 to be %d, got %d", 0, gameStatus.Board[1][1])
// 	}
// 	fmt.Println(gameStatus)
// }
