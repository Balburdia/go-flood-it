package main

import (
	"github.com/nsf/termbox-go"
)

// The different colors that can appear on the game board.
var colorMapping = []termbox.Attribute{
	termbox.ColorRed,
	termbox.ColorBlue,
	termbox.ColorGreen,
	termbox.ColorDarkGray,
	termbox.ColorMagenta,
	termbox.ColorYellow,
	termbox.ColorCyan,
	termbox.ColorLightGray,
}

// Board represents the game board.
type Board [][]int

type Point struct {
	X, Y int
}

func drawBoard(board Board) {
	// Clear the screen.
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y, row := range board {
		for x, cell := range row {
			// Draw the cell with two spaces to look better in the terminal
			termbox.SetCell(2*x, y, ' ', termbox.ColorDefault, colorMapping[cell])
			termbox.SetCell(2*x+1, y, ' ', termbox.ColorDefault, colorMapping[cell])
		}
	}

	// Refresh the screen to show the changes.
	termbox.Flush()
}

func floodFill(board Board, color int) {
	startColor := board[0][0]
	if startColor == color {
		// The start block is already the target color.
		return
	}

	// Create a visited 2D slice to keep track of which cells we've already visited.
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	// Start the BFS from the top-left cell.
	bfs(board, 0, 0, startColor, color, visited)
}

func bfs(board Board, x, y, startColor, newColor int, visited [][]bool) {
	// Initialize a queue with the starting point.
	queue := []Point{{X: x, Y: y}}

	// Iterate until the queue is empty.
	for len(queue) > 0 {
		// Pop a point from the front of the queue.
		point := queue[0]
		queue = queue[1:]

		// Check if this point has already been visited or if it's not the start color.
		if visited[point.Y][point.X] || board[point.Y][point.X] != startColor {
			continue
		}

		// Mark the point as visited and change its color.
		visited[point.Y][point.X] = true
		board[point.Y][point.X] = newColor

		// Add all valid neighboring points to the queue.
		if point.X > 0 {
			queue = append(queue, Point{X: point.X - 1, Y: point.Y})
		}
		if point.X < len(board[0])-1 {
			queue = append(queue, Point{X: point.X + 1, Y: point.Y})
		}
		if point.Y > 0 {
			queue = append(queue, Point{X: point.X, Y: point.Y - 1})
		}
		if point.Y < len(board)-1 {
			queue = append(queue, Point{X: point.X, Y: point.Y + 1})
		}
	}
}

func isGameOver(board Board) bool {
	// Get the color of the first cell.
	color := board[0][0]

	// Check all cells. If we find a cell that is not the same color, return false.
	for _, row := range board {
		for _, cell := range row {
			if cell != color {
				return false
			}
		}
	}

	// If we've checked all cells and they're all the same color, the game is over.
	return true
}
