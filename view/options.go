package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	statusText    *widget.Label
	restartButton *widget.Button
	playerToggle  *widget.RadioGroup
)

func NewOptions() *fyne.Container {
	statusText = widget.NewLabel("Player X's turn")
	restartButton = widget.NewButton("Restart", func() {
		gameStatus = control.Restart()
		updateUI()
	})
	playerToggle = widget.NewRadioGroup([]string{"Play as O"}, func(s string) {
		if s == "Play as O" {
			gameStatus = control.SwitchPlayer(true)
		} else {
			gameStatus = control.SwitchPlayer(false)
		}
		updateUI()
	})
	return container.NewVBox(statusText, restartButton, playerToggle)
}

func UpdateControlsUI(gameStatus string) {
	statusText.SetText(gameStatus)
}
