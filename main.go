package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	var (
		size   = flag.Int("size", 10, "Size of the game board.")
		colors = flag.Int("colors", 3, "Number of colors in the game.")
		seed   = flag.Int64("seed", time.Now().UnixNano(), "Seed number used in the random generator.")
		sleep  = flag.Duration("sleep", 500*time.Millisecond, "Time to sleep between steps.")
		noDraw = flag.Bool("nodraw", false, "Whether to draw the game board.")
	)

	// Parse command-line arguments.
	flag.Parse()

	if *colors < 3 || *colors > 8 {
		fmt.Println("Error: Number of colors must be between 3 and 8.")
		return
	}

	if *size <= 0 {
		fmt.Println("Error: Size must be a positive integer.")
		return
	}

	// Seed the random number generator.
	r := rand.New(rand.NewSource(*seed))

	// Initialize the game board.
	board := make(Board, *size)
	for i := range board {
		board[i] = make([]int, *size)
		for j := range board[i] {
			board[i][j] = r.Intn(*colors)
		}
	}

	// Initialize the UI.
	if !*noDraw {
		err := termbox.Init()
		if err != nil {
			panic(err)
		}
	}

	// Define the instance of the Algorithm used.
	// TODO: Replace the example bellow with your implementation
	// --------------------------------EXAMPLE-------------------------------------
	algorithm := &IncrementingColorAlgorithm{
		currentColor: board[0][0],
		numColors:    *colors,
	}
	// --------------------------------EXAMPLE-------------------------------------

	// Game loop.
	steps := 0
	for {
		// Draw the board.
		if !*noDraw {
			drawBoard(board)
		}

		// Get the next move from the algorithm
		color := algorithm.NextMove(board)

		// Apply the move to the board.
		floodFill(board, color)

		// Increase step counter.
		steps++

		// Check if the game is over.
		if isGameOver(board) {
			// Game is over, draw it one last time.
			if !*noDraw {
				drawBoard(board)
				time.Sleep(2 * time.Second)
			}
			break
		}

		if !*noDraw {
			time.Sleep(*sleep)
		}

	}

	// Closes the termbox. No defer was used to ensure that the print will
	// do its job without any even uglier solutions than this one.
	if !*noDraw {
		termbox.Close()
	}
	// Print the number of steps it took to finish the game.
	fmt.Printf("The game finished in %d steps.\n", steps)

}
