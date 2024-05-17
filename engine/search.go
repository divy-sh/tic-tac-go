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
		score := -negamax(newPos, 5, -math.MaxFloat64, math.MaxFloat64, false)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}

func negamax(pos *chess.Position, depth int, alpha, beta float64, isMaximizingPlayer bool) float64 {
	if depth == 0 {
		return eval(pos)
	}
	bestScore := -math.MaxFloat64
	moves := pos.ValidMoves()
	// Iterate over all moves
	for _, move := range moves {
		newPos := pos.Update(move)
		score := -negamax(newPos, depth-1, -beta, -alpha, !isMaximizingPlayer)
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
