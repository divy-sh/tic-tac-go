package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/divy-sh/tic-tac-go/game"
)

var (
	board         game.Game
	buttons       [][]*widget.Button
	statusText    *widget.Label
	restartButton *widget.Button
	playerToggle  *widget.RadioGroup
	isO           bool
)

func main() {
	a := app.New()
	w := a.NewWindow("Tic Tac Toe")
	board = game.NewGame(3)
	buttons = make([][]*widget.Button, 3)
	grid := container.NewAdaptiveGrid(3)

	for i := range buttons {
		buttons[i] = make([]*widget.Button, 3)
		for j := range buttons[i] {
			i, j := i, j // Capture i, j for each button
			buttons[i][j] = widget.NewButton(" ", func() {
				buttonClicked(i, j)
			})
			grid.Add(buttons[i][j])
		}
	}

	statusText = widget.NewLabel("Player X's turn")
	restartButton = widget.NewButton("Restart", func() {
		restartGame()
	})
	playerToggle = widget.NewRadioGroup([]string{"Play as X", "Play as Y"}, func(s string) {
		if s == "Play as X" {
			isO = false
		} else {
			isO = true
		}
		restartGame()
	})
	playerToggle.Selected = "Play as X"
	content := container.NewVSplit(grid, container.NewVBox(statusText, restartButton, playerToggle))

	w.SetContent(container.NewStack(content))
	w.ShowAndRun()
}

func restartGame() {
	board = game.NewGame(3)
	for i := range buttons {
		for j := range buttons[i] {
			buttons[i][j].SetText(" ")
			buttons[i][j].Enable()
		}
	}
	if isO {
		move := Eval(board)
		board, _ = board.PushMove(*move)
		updateUI()
	}
}

func buttonClicked(x, y int) {
	if board.IsGameOver() {
		return
	}

	if board.Board[x][y] != 0 {
		statusText.SetText("Invalid move, try again")
		return
	}

	newBoard, err := board.Move(x, y)
	if err != nil {
		statusText.SetText("Invalid move, try again")
		return
	}

	board = newBoard
	updateUI()

	if board.IsGameOver() {
		statusText.SetText(board.PrintGameStatus())
		return
	}

	move := Eval(board)
	board, _ = board.PushMove(*move)
	updateUI()

	if board.IsGameOver() {
		statusText.SetText(board.PrintGameStatus())
		return
	}
}

func updateUI() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch board.Board[i][j] {
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

	if board.Player == 1 {
		statusText.SetText("Player X's turn")
	} else {
		statusText.SetText("Player O's turn")
	}
}
