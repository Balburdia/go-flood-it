package main

// Algorithm is the interface that all Flood-It solving algorithms must implement.
type Algorithm interface {
	NextMove(Board) int
}

// --------------------------------EXAMPLE-------------------------------------
// TODO: Replace the example implementation bellow with yours.
type IncrementingColorAlgorithm struct {
	currentColor int
	numColors    int
}

func (a *IncrementingColorAlgorithm) NextMove(board Board) int {
	a.currentColor = (a.currentColor + 1) % a.numColors
	return a.currentColor
}

// --------------------------------EXAMPLE-------------------------------------
