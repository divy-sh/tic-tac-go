package controller

import (
	"github.com/divy-sh/tic-tac-go/engine"
	"github.com/divy-sh/tic-tac-go/game"
)

type Controller struct {
	isO   bool
	board game.Game
}

func NewController() *Controller {
	return &Controller{
		isO:   false,
		board: game.NewGame(3),
	}
}

func (c *Controller) Restart() *GameStatus {
	c.board = game.NewGame(3)
	if c.isO {
		move := engine.Eval(c.board, c.isO)
		c.board, _ = c.board.PushMove(*move)
	}
	return c.GetGameStatus()
}

func (c *Controller) Move(x, y int) (*GameStatus, error) {
	if c.board.IsGameOver() {
		return c.GetGameStatus(), nil
	}

	newBoard, err := c.board.Move(x, y)
	if err != nil {
		return nil, err
	}

	c.board = newBoard

	if c.board.IsGameOver() {
		return c.GetGameStatus(), nil
	}

	move := engine.Eval(c.board, c.isO)
	c.board, _ = c.board.PushMove(*move)
	return c.GetGameStatus(), nil
}

func (c *Controller) SwitchPlayer(isO bool) *GameStatus {
	c.isO = isO
	return c.Restart()
}

func (c *Controller) GetGameStatus() *GameStatus {
	return &GameStatus{
		Board:      c.board.Board,
		GameStatus: c.board.PrintGameStatus(),
	}
}
