package engine

import (
	"math"

	"github.com/notnil/chess"
)

func search(depth int, pos *chess.Position) *chess.Move {
	bestScore := -math.MaxFloat32
	var bestMove *chess.Move
	moves := pos.ValidMoves()
	for _, move := range moves {
		newPos := pos.Update(move)
		score := -negamax(newPos, 5, -math.MaxInt32, math.MaxInt32, false)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}

func negamax(pos *chess.Position, depth int, alpha, beta int, isMaximizingPlayer bool) float64 {
	if depth == 0 {
		return eval(pos)
	}
	bestScore := -math.MaxInt32
	moves := board.ValidMoves()
	orderedMoves := orderMoves(moves, board)
	// Iterate over all moves
	for _, move := range orderedMoves {
		newBoard := board.Clone()
		newBoard.Position().Update(move)
		score := -negamax(newBoard, depth-1, -beta, -alpha, !isMaximizingPlayer)
		// Update the best score and alpha value
		bestScore = max(bestScore, score)
		alpha = max(alpha, score)
		// Pruning condition
		if alpha >= beta {
			break
		}
	}
	return bestScore
}
