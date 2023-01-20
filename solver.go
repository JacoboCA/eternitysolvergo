package main

// functions to solve the puzzle with several strategies and instructions

func solveDummy(puzzle *PuzzleState, depth int, strategy int, queue *[]*PuzzleState) bool {
	// this should be thread-safe as it has its own data and queue

	// set up working environment
	base := puzzle.base
	pieces := puzzle.pieces
	colours := base.getColours()
	bordersByColour := make([][]*PieceState, 0, len(colours))
	piecesByColour := make([][]*PieceState, 0, len(colours))
	for k := 0; k < len(colours); k++ {
		bordersByColour[k] = make([]*PieceState, 0)
		piecesByColour[k] = make([]*PieceState, 0)
	}
	corners := make([]*PieceState, 0, 4)

	for i := 0; i < base.nPieces; i++ {
		if piece := pieces[i]; !piece.inUse {
			if piece.base.isCorner {
				corners = append(corners, piece)
			} else if piece.base.isBorder {
				bordersByColour[piece.north] = append(bordersByColour[piece.north], piece)
				bordersByColour[piece.east] = append(bordersByColour[piece.east], piece)
				bordersByColour[piece.south] = append(bordersByColour[piece.south], piece)
				bordersByColour[piece.west] = append(bordersByColour[piece.west], piece)
			} else if piece.base.isMiddle {
				piecesByColour[piece.north] = append(piecesByColour[piece.north], piece)
				piecesByColour[piece.east] = append(piecesByColour[piece.east], piece)
				piecesByColour[piece.south] = append(piecesByColour[piece.south], piece)
				piecesByColour[piece.west] = append(piecesByColour[piece.west], piece)
			}
		}
	}

	// we have in piecesByColour and bordersByColour a set of arrays of each colour of available pieces
	// we can now loop adding pieces to the board based on strategy knowing the initial state (does not change) and knowing the available pieces

	return true
}
