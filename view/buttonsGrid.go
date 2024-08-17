package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	buttons [][]*widget.Button
)

func NewButtonsGrid(size int) *fyne.Container {
	grid := container.NewAdaptiveGrid(size)
	buttons = make([][]*widget.Button, size)

	for i := range buttons {
		buttons[i] = make([]*widget.Button, size)
		for j := range buttons[i] {
			i, j := i, j // Capture i, j for each button
			buttons[i][j] = widget.NewButton(" ", func() {
				buttonClicked(i, j)
			})
			grid.Add(buttons[i][j])
		}
	}
	return grid
}

func buttonClicked(x, y int) {
	status, err := control.Move(x, y)
	if err != nil {
		gameStatus.GameStatus = err.Error()
	} else {
		gameStatus = status
	}
	updateUI()
}

func UpdateButtonsUI(board [][]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch board[i][j] {
			case 1:
				buttons[i][j].SetText("❌")
				buttons[i][j].Disable()
			case -1:
				buttons[i][j].SetText("⭕️")
				buttons[i][j].Disable()
			case 0:
				buttons[i][j].SetText(" ")
				buttons[i][j].Enable()
			}
		}
	}
}
