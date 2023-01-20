package main

type Coordinator struct {
	workers []*Worker
	// to-do ids
}

func (c *Coordinator) newWorker(start *PuzzleState, depth int, strategy int) int {
	var queue []*PuzzleState = make([]*PuzzleState, 10)
	if c.workers == nil {
		c.workers = make([]*Worker, 0, 1)
	}
	c.workers = append(c.workers, newWorker(start, depth, strategy, &queue))
	return len(c.workers) - 1
}

func (c *Coordinator) run(id int) bool {
	return c.workers[id].run()
}
