package engine

import "github.com/notnil/chess"

var (
	pieceVal = map[chess.PieceType]float64{
		chess.Rook: 500, chess.Pawn: 100, chess.Bishop: 330, chess.Knight: 320, chess.Queen: 900, chess.King: 10000,
	}

	pawnTable = [64]float64{
		0, 0, 0, 0, 0, 0, 0, 0,
		50, 50, 50, 50, 50, 50, 50, 50,
		10, 10, 20, 30, 30, 20, 10, 10,
		5, 5, 10, 25, 25, 10, 5, 5,
		0, 0, 0, 20, 20, 0, 0, 0,
		5, -5, -10, 0, 0, -10, -5, 5,
		5, 10, 10, -20, -20, 10, 10, 5,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	revPawnTable = reversed(pawnTable)

	knightTable = [64]float64{
		-50, -40, -30, -30, -30, -30, -40, -50,
		-40, -20, 0, 0, 0, 0, -20, -40,
		-30, 0, 10, 15, 15, 10, 0, -30,
		-30, 5, 15, 20, 20, 15, 5, -30,
		-30, 0, 15, 20, 20, 15, 0, -30,
		-30, 5, 10, 15, 15, 10, 5, -30,
		-40, -20, 0, 5, 5, 0, -20, -40,
		-50, -40, -30, -30, -30, -30, -40, -50,
	}

	revKnightTable = reversed(knightTable)

	bishopTable = [64]float64{
		-20, -10, -10, -10, -10, -10, -10, -20,
		-10, 0, 0, 0, 0, 0, 0, -10,
		-10, 0, 5, 10, 10, 5, 0, -10,
		-10, 5, 5, 10, 10, 5, 5, -10,
		-10, 0, 10, 10, 10, 10, 0, -10,
		-10, 10, 10, 10, 10, 10, 10, -10,
		-10, 5, 0, 0, 0, 0, 5, -10,
		-20, -10, -10, -10, -10, -10, -10, -20,
	}

	revBishopTable = reversed(bishopTable)

	rookTable = [64]float64{
		0, 0, 0, 0, 0, 0, 0, 0,
		5, 10, 10, 10, 10, 10, 10, 5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		0, 0, 0, 5, 5, 0, 0, 0,
	}

	revRookTable = reversed(rookTable)

	queenTable = [64]float64{
		-20, -10, -10, -5, -5, -10, -10, -20,
		-10, 0, 0, 0, 0, 0, 0, -10,
		-10, 0, 5, 5, 5, 5, 0, -10,
		-5, 0, 5, 5, 5, 5, 0, -5,
		0, 0, 5, 5, 5, 5, 0, -5,
		-10, 5, 5, 5, 5, 5, 0, -10,
		-10, 0, 5, 0, 0, 0, 0, -10,
		-20, -10, -10, -5, -5, -10, -10, -20,
	}

	revQueenTable = reversed(queenTable)

	kingMiddleGameTable = [64]float64{
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-20, -30, -30, -40, -40, -30, -30, -20,
		-10, -20, -20, -20, -20, -20, -20, -10,
		20, 20, 0, 0, 0, 0, 20, 20,
		20, 30, 10, 0, 0, 10, 30, 20,
	}

	revKingMiddleGameTable = reversed(kingMiddleGameTable)

	kingEndGameTable = [64]float64{
		-50, -40, -30, -20, -20, -30, -40, -50,
		-30, -20, -10, 00, 00, -10, -20, -30,
		-30, -10, 20, 30, 30, 20, -10, -30,
		-30, -10, 30, 40, 40, 30, -10, -30,
		-30, -10, 30, 40, 40, 30, -10, -30,
		-30, -10, 20, 30, 30, 20, -10, -30,
		-30, -30, 00, 00, 00, 00, -30, -30,
		-50, -30, -30, -30, -30, -30, -30, -50,
	}

	revKingEndGameTable = [64]float64{
		-50, -30, -30, -30, -30, -30, -30, -50,
		-30, -30, 00, 00, 00, 00, -30, -30,
		-30, -10, 20, 30, 30, 20, -10, -30,
		-30, -10, 30, 40, 40, 30, -10, -30,
		-30, -10, 30, 40, 40, 30, -10, -30,
		-30, -10, 20, 30, 30, 20, -10, -30,
		-30, -20, -10, 00, 00, -10, -20, -30,
		-50, -40, -30, -20, -20, -30, -40, -50,
	}
)

// returns a basic eval of the position
func eval(pos *chess.Position) float64 {
	var value float64
	var materialVal float64
	var positionalVal float64
	var mobilityVal float64
	centerControlVal := evaluateCenterControl(pos)
	for square, piece := range pos.Board().SquareMap() {
		if piece.Color() == chess.White {
			materialVal += pieceVal[piece.Type()]
			positionalVal += getPieceSquareTable(pos, &square, &piece)
		} else {
			materialVal -= pieceVal[piece.Type()]
			positionalVal -= getPieceSquareTable(pos, &square, &piece)
		}
	}
	value = materialVal + positionalVal + mobilityVal + centerControlVal
	gamePhase := getGamePhase(pos)
	value = value*gamePhase + (1-gamePhase)*materialVal
	return value
}

func getPieceSquareTable(pos *chess.Position, square *chess.Square, piece *chess.Piece) float64 {
	sq := int(*square)
	if piece.Color() == chess.White {
		switch piece.Type() {
		case chess.Pawn:
			return pawnTable[sq]
		case chess.Knight:
			return knightTable[sq]
		case chess.Bishop:
			return bishopTable[sq]
		case chess.Rook:
			return rookTable[sq]
		case chess.Queen:
			return queenTable[sq]
		case chess.King:
			gamePhase := getGamePhase(pos)
			return gamePhase*kingMiddleGameTable[sq] + (1-gamePhase)*kingEndGameTable[sq]
		}
	} else {
		switch piece.Type() {
		case chess.Pawn:
			return revPawnTable[sq]
		case chess.Knight:
			return revKnightTable[sq]
		case chess.Bishop:
			return revBishopTable[sq]
		case chess.Rook:
			return revRookTable[sq]
		case chess.Queen:
			return revQueenTable[sq]
		case chess.King:
			gamePhase := getGamePhase(pos)
			return gamePhase*revKingMiddleGameTable[sq] + (1-gamePhase)*revKingEndGameTable[sq]
		}
	}
	return 0
}

// evaluateCenterControl evaluates the control of the center squares (e4, e5, d4, d5)
func evaluateCenterControl(pos *chess.Position) float64 {
	var score float64
	centerSquares := []chess.Square{chess.E4, chess.E5, chess.D4, chess.D5}

	for _, square := range centerSquares {
		if pos.Board().Piece(square) != chess.NoPiece {
			if pos.Board().Piece(square).Color() == chess.White {
				score += pieceVal[pos.Board().Piece(square).Type()]
			} else {
				score -= pieceVal[pos.Board().Piece(square).Type()]
			}
		}
	}
	return score
}

func getGamePhase(pos *chess.Position) float64 {
	totalPieces := float64(len(pos.Board().SquareMap()))
	maxPieces := 32.0
	gamePhase := totalPieces / maxPieces
	return float64(gamePhase)
}

func reversed(pst [64]float64) [64]float64 {
	reversed := [64]float64{}
	for i, value := range pst {
		reversed[63-i] = value
	}
	return reversed
}
