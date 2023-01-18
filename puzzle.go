package main

import "fmt"

// Functions and classes related to puzzle pieces and board

// puzzle board , only squared boards accepted
type Puzzle struct {
	pieces  []*Piece
	size    uint8
	nPieces int
}

func newPuzzle(pieces []*Piece, size uint8, nPieces int) *Puzzle {
	puzzle := new(Puzzle)
	puzzle.size = size

	puzzle.nPieces = nPieces
	puzzle.pieces = pieces

	return puzzle
}

// index i,j starts top left, i does rows
type PuzzleState struct {
	base   *Puzzle
	pieces []*PieceState
	board  [][]*PieceState
}

func newPuzzleStateEmpty(base *Puzzle) *PuzzleState {
	state := new(PuzzleState)
	state.base = base
	state.pieces = make([]*PieceState, state.base.nPieces)

	board := make([][]*PieceState, state.base.size)
	for i := 0; i < int(state.base.size); i++ {
		board[i] = make([]*PieceState, state.base.size)
	}

	for i := 0; i < int(state.base.nPieces); i++ {
		state.pieces[i] = newPieceState(state.base.pieces[i], 0, -1, -1)
	}

	state.board = board
	return state
}

func newPuzzleState(base *Puzzle, pieces []*PieceState) *PuzzleState {
	state := new(PuzzleState)
	state.base = base
	state.pieces = pieces

	board := make([][]*PieceState, state.base.size)
	for i := 0; i < int(state.base.size); i++ {
		board[i] = make([]*PieceState, state.base.size)
	}

	for i := 0; i < int(state.base.nPieces); i++ {
		p := state.pieces[i]
		fmt.Println(i, p)
		if p.x > -1 {
			board[p.x][p.y] = p
		}
	}

	state.board = board
	return state
}

// object piece with 4 colours in the cardinal
type Piece struct {
	id      int
	raw     [4]byte
	colours [7]byte
}

// rotation and cardinals are part of the state
type PieceState struct {
	base     *Piece
	north    byte
	east     byte
	south    byte
	west     byte
	rotation byte
	x, y     int8
}

// create for readability but cannot be changed and expected to work. E.g. cannot make a triangle/hexagon by changing to 6 as the position array is set to 7 and would overflow
const PieceMaxRotations int = 4

func (p *PieceState) rotateClockunwise() {
	// circular 0->1->2->3
	p.rotation = (p.rotation + 1) % 4 // 4 rotations
	p.north = p.base.colours[0+p.rotation]
	p.east = p.base.colours[1+p.rotation]
	p.south = p.base.colours[2+p.rotation]
	p.west = p.base.colours[3+p.rotation]
}

func (p *PieceState) rotateClockwise() {
	// circular 0->1->2->3
	p.rotation = (p.rotation + 6) % 4 // 4 rotations
	p.north = p.base.colours[0+p.rotation]
	p.east = p.base.colours[1+p.rotation]
	p.south = p.base.colours[2+p.rotation]
	p.west = p.base.colours[3+p.rotation]
}

func (p *PieceState) resetRaw() *PieceState {

	p.base.colours[0] = byte(255)
	p.base.colours[1] = byte(255)
	p.base.colours[2] = byte(255)
	p.base.colours[3] = byte(255)
	p.base.colours[4] = byte(255)
	p.base.colours[5] = byte(255)
	p.base.colours[6] = byte(255)

	p.base.colours[0+p.rotation] = p.north
	p.base.colours[1+p.rotation] = p.east
	p.base.colours[2+p.rotation] = p.south
	p.base.colours[3+p.rotation] = p.west

	for i := 0; i < 7; i++ {
		if p.base.colours[i] == 255 {
			p.base.colours[i] = p.base.colours[(i+PieceMaxRotations)%7]
		}
	}

	p.base.raw[0] = p.base.colours[0]
	p.base.raw[1] = p.base.colours[1]
	p.base.raw[2] = p.base.colours[2]
	p.base.raw[3] = p.base.colours[3]

	return p
}

// create new piece with the colours
func newPieceState(p *Piece, rotation byte, x, y int8) *PieceState {

	return &PieceState{
		base:     p,
		north:    p.colours[0+rotation],
		east:     p.colours[1+rotation],
		south:    p.colours[2+rotation],
		west:     p.colours[3+rotation],
		rotation: rotation,
		x:        x,
		y:        y,
	}
}

func newPiece(id int, north, east, south, west byte) *Piece {
	return &Piece{id: id,
		raw:     [4]byte{north, east, south, west},
		colours: [7]byte{north, east, south, west, north, east, south},
	}
}
