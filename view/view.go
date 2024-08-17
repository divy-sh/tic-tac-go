package view

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/divy-sh/tic-tac-go/controller"
)

var (
	control    *controller.Controller
	gameStatus *controller.GameStatus
)

func NewView(c *controller.Controller) {
	control = c
	gameStatus = control.GetGameStatus()

	a := app.New()
	w := a.NewWindow("Tic Tac Toe")

	content := container.NewVSplit(NewButtonsGrid(3), NewOptions())

	w.SetContent(container.NewStack(content))
	w.ShowAndRun()
}

func updateUI() {
	UpdateButtonsUI(gameStatus.Board)
	UpdateControlsUI(gameStatus.GameStatus)
}
