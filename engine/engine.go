package engine

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
)

func GenMoveIterative(seconds int, pos *chess.Position) *chess.Move {
	defer timer("genMoveIterative")()
	done := make(chan bool)
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		done <- true
	}()
	depth := 2
	var move *chess.Move = nil
	for {
		select {
		case <-done:
			fmt.Printf("depth reached - %d\n", depth)
			return move
		default:
			depth++
			move = GenMove(depth, pos)
		}
	}
}

func GenMove(depth int, pos *chess.Position) *chess.Move {
	defer timer("genMove")()
	move := search(depth, pos)
	return move
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
