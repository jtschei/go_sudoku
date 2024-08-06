# Sudoku

This Go project is a Sudoku puzzle generator and solver.

## Using

```
Usage of sudoku:
  -blanks int
        # of blank values in puzzle (default 60)
  -puzzle
        user will provide puzzle to stdin
  -repeat int
        # of times to repeat puzzle solve
  -swaps int
        # of randomization swaps (default 40)
```

## Building

`go build`

## Testing

```
cd sudoku
go test
```

## Example Puzzles

- `cat examples/puzzle.stdin | sudoku -puzzle`
- `cat examples/puzzle2.stdin | sudoku -puzzle`
- `cat examples/puzzle3.stdin | sudoku -puzzle`
