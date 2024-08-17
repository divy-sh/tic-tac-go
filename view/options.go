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
		control.Restart()
	})
	playerToggle = widget.NewRadioGroup([]string{"Play as X", "Play as Y"}, func(s string) {

	})
	playerToggle.Selected = "Play as X"
	return container.NewVBox(statusText, restartButton, playerToggle)
}

func UpdateControlsUI(gameStatus string) {
	statusText.SetText(gameStatus)
}
