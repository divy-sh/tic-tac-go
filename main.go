package main

import (
	"github.com/divy-sh/tic-tac-go/controller"
	"github.com/divy-sh/tic-tac-go/view"
)

func main() {
	controller := controller.NewController()
	view.NewView(controller)
}
