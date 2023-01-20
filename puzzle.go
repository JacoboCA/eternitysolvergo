package main

import "fmt"

// Functions and classes related to puzzle pieces and board

// create for readability but cannot be changed and expected to work. E.g. cannot make a triangle/hexagon by changing to 6 as the position array is set to 7 and would overflow
const PieceMaxRotations int = 4
const BorderColour int8 = 0

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
	id       int
	raw      [4]int8
	colours  [7]int8
	isBorder bool
	isCorner bool
	isMiddle bool
}

// rotation and cardinals are part of the state
type PieceState struct {
	base     *Piece
	north    int8
	east     int8
	south    int8
	west     int8
	rotation int8
	x, y     int8
	inUse    bool
}

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

	p.base.colours[0] = -1
	p.base.colours[1] = -1
	p.base.colours[2] = -1
	p.base.colours[3] = -1
	p.base.colours[4] = -1
	p.base.colours[5] = -1
	p.base.colours[6] = -1

	p.base.colours[0+p.rotation] = p.north
	p.base.colours[1+p.rotation] = p.east
	p.base.colours[2+p.rotation] = p.south
	p.base.colours[3+p.rotation] = p.west

	for i := 0; i < 7; i++ {
		if p.base.colours[i] == -1 {
			p.base.colours[i] = p.base.colours[(i+PieceMaxRotations)%7]
		}
	}

	p.base.raw[0] = p.base.colours[0]
	p.base.raw[1] = p.base.colours[1]
	p.base.raw[2] = p.base.colours[2]
	p.base.raw[3] = p.base.colours[3]

	isBorder := p.north == BorderColour || p.east == BorderColour || p.south == BorderColour || p.west == BorderColour
	isMiddle := !isBorder
	isCorner := false
	if isBorder {
		isCorner = p.north == p.east && p.north == BorderColour || p.east == p.south && p.east == BorderColour || p.south == p.west && p.south == BorderColour || p.west == p.north && p.west == BorderColour
		isBorder = !isCorner
	}

	p.base.isBorder = isBorder
	p.base.isCorner = isCorner
	p.base.isMiddle = isMiddle
	return p
}

// create new piece with the colours
func newPieceState(p *Piece, rotation int8, x, y int8) *PieceState {
	inUse := x >= 0
	return &PieceState{
		base:     p,
		north:    p.colours[0+rotation],
		east:     p.colours[1+rotation],
		south:    p.colours[2+rotation],
		west:     p.colours[3+rotation],
		rotation: rotation,
		x:        x,
		y:        y,
		inUse:    inUse,
	}
}

func newPiece(id int, north, east, south, west int8) *Piece {
	isBorder := north == BorderColour || east == BorderColour || south == BorderColour || west == BorderColour
	isMiddle := !isBorder
	isCorner := false
	if isBorder {
		isCorner = north == east && north == BorderColour || east == south && east == BorderColour || south == west && south == BorderColour || west == north && west == BorderColour
		isBorder = !isCorner
	}

	return &Piece{id: id,
		raw:      [4]int8{north, east, south, west},
		colours:  [7]int8{north, east, south, west, north, east, south},
		isBorder: isBorder,
		isCorner: isCorner,
		isMiddle: isMiddle,
	}
}

func (p *PuzzleState) isValid() bool {
	for i := 0; i < int(p.base.size)-1; i++ { // always has a next col and next row
		for j := 0; j < int(p.base.size)-1; j++ {
			if p.board[i][j] != nil && p.board[i+1][j] != nil && p.board[i][j].east != p.board[i+1][j].west {
				return false
			}
			if p.board[i][j] != nil && p.board[i][j+1] != nil && p.board[i][j].south != p.board[i][j+1].north {
				return false
			}
		}
	}
	return true
}

func (p *PuzzleState) clone() *PuzzleState {

	var pieces []*PieceState = make([]*PieceState, 0, p.base.nPieces)
	for i := 0; i < p.base.nPieces; i++ {
		piece := p.pieces[i]
		pieces = append(pieces, newPieceState(piece.base, piece.rotation, piece.x, piece.y))
	}
	res := newPuzzleState(p.base, pieces)

	return res
}

func (p *Puzzle) getColours() []int8 {
	var haystack map[int8]struct{}

	haystack = make(map[int8]struct{})

	for i := 0; i < p.nPieces; i++ {
		piece := p.pieces[i]
		for k := 0; k < PieceMaxRotations; k++ {
			haystack[piece.raw[k]] = struct{}{}
		}
	}

	keys := make([]int8, 0, len(haystack))
	for k := range haystack {
		keys = append(keys, k)
	}

	return keys
}
