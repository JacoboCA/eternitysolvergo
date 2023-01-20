package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

// run eternitysolver
func main() {
	// set seed at beginning to make sure it is reseeded each time
	rand.Seed(time.Now().UnixNano())
	if true {
		leader()
	} else {
		coordinator := new(Coordinator)
		testCoordinator(coordinator)
	}
}

func loadPuzzle(path string) *Puzzle {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileReader := bufio.NewScanner(file)
	fileReader.Split(bufio.ScanLines) // already by default

	var pieces []*Piece
	pieces = make([]*Piece, 0, 256)

	for i := 0; fileReader.Scan(); i++ {
		var id int
		var north, east, south, west int8

		fmt.Sscanf(fileReader.Text(), "%d %d %d %d %d", &id, &north, &east, &south, &west)

		pieces = append(pieces, newPiece(id, north, east, south, west))
	}

	nPieces := len(pieces)
	var size uint8 = uint8(math.Floor(math.Sqrt(float64(nPieces))))

	var puzzle *Puzzle = newPuzzle(pieces, size, nPieces)

	return puzzle
}

func savePuzzle(puzzle *Puzzle, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buff := bufio.NewWriter(file)

	for i := 0; i < puzzle.nPieces; i++ {
		p := puzzle.pieces[i]
		_, err = fmt.Fprintf(buff, "%d %d %d %d %d\n", p.id, p.raw[0], p.raw[1], p.raw[2], p.raw[3])
		if err != nil {
			panic(err)
		}
	}

	buff.Flush()

}

func savePuzzleState(state *PuzzleState, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buff := bufio.NewWriter(file)

	for i := 0; i < int(state.base.nPieces); i++ {
		piece := state.pieces[i]

		_, err = fmt.Fprintf(buff, "%d %d %d %d\n", piece.base.id, piece.rotation, piece.x, piece.y)
		if err != nil {
			panic(err)
		}

	}

	buff.Flush()

}

func loadPuzzleState(path string, puzzle *Puzzle) *PuzzleState {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileReader := bufio.NewScanner(file)
	fileReader.Split(bufio.ScanLines) // already by default

	var pieces []*PieceState
	pieces = make([]*PieceState, 0, puzzle.nPieces)

	for i := 0; fileReader.Scan(); i++ {
		var id int
		var rotation int8
		var x, y int8

		fmt.Sscanf(fileReader.Text(), "%d %d %d %d", &id, &rotation, &x, &y)

		pieces = append(pieces, newPieceState(puzzle.pieces[id], rotation, x, y).resetRaw())
	}

	// nPieces := len(pieces)
	// var size uint8 = uint8(math.Floor(math.Sqrt(float64(nPieces))))

	var state *PuzzleState = newPuzzleState(puzzle, pieces)

	return state
}

// print state in 3x3 cells for each piece
func printPuzzleState(puzzle *PuzzleState) {
	var size int = int(puzzle.base.size)

	// get separating line
	line := "-"
	for i := 0; i < int(size); i++ {
		line = line + "----"
	}

	// print first line
	fmt.Println(line)
	for i := 0; i < size; i++ {
		line1 := "|"
		line2 := "|"
		line3 := "|"

		for j := 0; j < size; j++ {
			p := puzzle.board[i][j]
			if p == nil {
				line1 = line1 + fmt.Sprintf("   |")
				line2 = line2 + fmt.Sprintf("   |")
				line3 = line3 + fmt.Sprintf("   |")
			} else {
				line1 = line1 + fmt.Sprintf(" %d |", p.north)
				line2 = line2 + fmt.Sprintf("%d %d|", p.west, p.east)
				line3 = line3 + fmt.Sprintf(" %d |", p.south)
			}
		}

		fmt.Println(line1)
		fmt.Println(line2)
		fmt.Println(line3)
		fmt.Println(line)
	}

}
