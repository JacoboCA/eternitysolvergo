# Eternity II Solver in Go
#### _A brute force attemp to solve the [Eternity II puzzle](https://en.wikipedia.org/wiki/Eternity_II_puzzle) in go._

Eternity solver is managed by a _leader_ who divides the problem for _workers_ to solve.
- leader: starts the puzzle solving, divides the task, manages and controls the progress, communicates the task to the workers
- worker: get a initial state and a depth and calculates all feasible solutions 

# Development
This is a personal project used for learning go.
