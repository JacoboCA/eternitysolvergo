package main

// this allows to do multithread in same computer and have a coordinator layer
type Worker struct {
	puzzle   *PuzzleState
	depth    int
	strategy int
	queue    *[]*PuzzleState
}

func newWorker(start *PuzzleState, depth int, strategy int, queue *[]*PuzzleState) *Worker {
	return &Worker{start, depth, strategy, queue}
}

func (w *Worker) run() bool {

	return solveDummy(w.puzzle, w.depth, w.strategy, w.queue)

	// 	return true
}
