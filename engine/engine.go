package engine

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
)

func GenMoveIterative(seconds int, game *chess.Game) *chess.Move {
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
			move = GenMove(depth, game)
		}
	}
}

func GenMove(depth int, game chess.Game) *chess.Move {
	defer timer("genMove")()
	move := search(depth, game)
	return move
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
