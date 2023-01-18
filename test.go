package main

import (
	"math/rand"
	"sort"
)

// test functions
func createNewPuzzle(size int, colours int, basepath string) *PuzzleState {
	nPieces := size * size
	pieces := []*Piece{}

	for i := 0; i < nPieces; i++ {
		pieces = append(pieces, newPiece(0, 0, 0, 0, 0))
	}
	puzzle := newPuzzle(pieces, uint8(size), nPieces)

	state := newPuzzleStateEmpty(puzzle)

	// add pieces to the board (currently all pieces are 0,0,0,0)
	for i := 0; i < nPieces; i++ {
		state.board[i/size][i%size] = state.pieces[i]
	}

	// random ids
	// crate arrayfor ids, shuffle, assign linearly
	ids := make([]int, nPieces)
	for i := 0; i < nPieces; i++ {
		ids[i] = i
	}

	rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })

	// shuffle colours and rotation, add id
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// i is rows, 0,0 is top left
			piece := state.board[i][j]

			piece.base.id = ids[i*size+j]
			piece.rotation = byte(rand.Intn(PieceMaxRotations))
			piece.x = int8(i)
			piece.y = int8(j)

			if i != 0 {
				piece.north = state.board[i-1][j].south
			}

			if i != size-1 {
				piece.south = byte(rand.Intn(colours) + 1)
			}

			if j != 0 {
				piece.west = state.board[i][j-1].east
			}

			if j != size-1 {
				piece.east = byte(rand.Intn(colours) + 1)
			}

			piece.resetRaw()
		}
	}

	sort.Slice(puzzle.pieces, func(i, j int) bool {
		return puzzle.pieces[i].id < puzzle.pieces[j].id
	})

	sort.Slice(state.pieces, func(i, j int) bool {
		return state.pieces[i].base.id < state.pieces[j].base.id
	})

	savePuzzle(puzzle, basepath+"puzzle.txt")
	savePuzzleState(state, basepath+"solution.txt")
	// export solution as state

	// export puzzle as base

	return state
}
