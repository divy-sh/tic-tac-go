package engine

import (
	"math"
	"sort"

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

func alphaBeta(depth int, game chess.Game, alpha float64, beta float64, isMax bool) (*chess.Move, float64) {
	if depth == 0 {
		return nil, eval(game)
	}
	bestValue := 0.0
	if isMax {
		bestValue = math.Inf(-1)
	} else {
		bestValue = math.Inf(1)
	}
	var bestMove *chess.Move
	moves := game.Position().ValidMoves()
	sort.Slice(moves, func(i int, j int) bool {
		iPriority := 0.0
		jPriority := 0.0
		// check if promotion move
		if moves[i].Promo() != chess.NoPiece.Type() {
			iPriority += pieceVal[moves[i].Promo()] / pieceVal[chess.Queen]
		}
		if moves[j].Promo() != chess.NoPiece.Type() {
			jPriority += pieceVal[moves[j].Promo()] / pieceVal[chess.Queen]
		}
		// check if capture move
		if game.Position().Board().Piece(moves[i].S2()) != chess.NoPiece {
			iPriority += pieceVal[game.Position().Board().Piece(moves[i].S2()).Type()] / pieceVal[chess.Queen]
		}
		if game.Position().Board().Piece(moves[j].S2()) != chess.NoPiece {
			iPriority += pieceVal[game.Position().Board().Piece(moves[j].S2()).Type()] / pieceVal[chess.Queen]
		}
		//check if check
		if moves[i].HasTag(chess.Check) {
			iPriority += 0.5
		}
		if moves[j].HasTag(chess.Check) {
			jPriority += 0.5
		}
		return iPriority > jPriority
	})
	for _, move := range moves {
		newGame := game.Clone()
		newGame.Position().Update(move)
		_, value := alphaBeta(depth-1, *newGame, alpha, beta, !isMax)
		if isMax {
			if bestValue <= value {
				bestValue = value
				bestMove = move
			}
			alpha = math.Max(alpha, bestValue)
		} else {
			if bestValue >= value {
				bestValue = value
				bestMove = move
			}
			beta = math.Min(beta, bestValue)
		}
		if beta <= alpha {
			return bestMove, bestValue
		}
	}
	return bestMove, bestValue
}
