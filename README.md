# go-flood-it
Template for testing heuristics to beat the game Flood It written in Golang.

## Adding a new heuristic/algorithm
In order to add new algorithm to beat the game you must make changes in two files:

- `algorithm.go`: Your algorithm must implement the interface `Algorithm` defined on this file and its method `NextMove`. `NextMove` receives the board and returns the decision for the next move (color). The example bellow implements an Algorithm that will just cycle through all the colors in the game sequencially.
```go
// Algorithm is the interface that all Flood-It solving algorithms must implement.
type Algorithm interface {
	NextMove(Board) int
}

// Example implementation
type IncrementingColorAlgorithm struct {
	currentColor int
	numColors    int
}

func (a *IncrementingColorAlgorithm) NextMove(board Board) int {
	a.currentColor = (a.currentColor + 1) % a.numColors
	return a.currentColor
}
```
- `main.go`: Inside the main method you just need to initialize your algorithm. The example bellow creates an instance of the `IncrementingColorAlgorithm` described above.
```go
func main() {
    // ...

    // Define the instance of the Algorithm used.
    algorithm := &IncrementingColorAlgorithm{
        currentColor: board[0][0],
        numColors:    *colors,
    }
    
    // Game loop.
    // ..
```

## Running the game
In order to run the game you just need to build your application with `go build` and execute the `go-flood-it` executable created by it. Bellow you will find some parameters to customize your run.

### Parameters

- `--size`: Defines the size of the board. The board is always a square of NxN, where N is the value defined in this parameter. Default: `10`.

- `--colors`: Number of different colors to use in the game. The number must be minimum 3 and maximum 8. Default: `3`.

- `--seed`: Seed number to be used in all random-based decisions in the game when creating its board. Default: `time.Now().UnixNano()`

- `--sleep`: Time to sleep between steps. Default: `500ms`. Usage: `--sleep 1s`, `--sleep 250ms`.

- `--nodraw`: If specified, the board will not be drawn and only the number of steps will be printed by the end of execution. This also removes the waiting between steps.
