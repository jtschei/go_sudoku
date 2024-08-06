package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	model "github.com/jtschei/sudoku"
)

/*
Reads lines of input from stdin and builds puzzle from it.
The format of puzzle should be same as format of puzzle when it is printed to string.
*/
func read_puzzle_from_stdin() *model.Puzzle {
	//read puzzle
	scanner := bufio.NewScanner(os.Stdin)
	puzzle_str := ""
	for scanner.Scan() {
		line := scanner.Text()
		puzzle_str = puzzle_str + line + "\n"
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	//build and return
	puzzle := model.Puzzle_From_String(puzzle_str)
	return puzzle
}

func main() {
	//arguments
	swap_cnt := flag.Int("swaps", 40, "# of randomization swaps")
	blank_cnt := flag.Int("blanks", 60, "# of blank values in puzzle")
	repeat_cnt := flag.Int("repeat", 0, "# of times to repeat puzzle solve")
	puzzle_flag := flag.Bool("puzzle", false, "user will provide puzzle to stdin")
	flag.Parse()

	var puzzle *model.Puzzle

	for repeat_idx := 0; repeat_idx <= *repeat_cnt; repeat_idx++ {

		if *puzzle_flag {
			//read user supplied puzzle
			puzzle = read_puzzle_from_stdin()
			fmt.Println("Read puzzle from stdin")
			//don't repeat user supplied
			if *repeat_cnt != 0 {
				fmt.Println("Warning: Ignoring repeat argument")
				*repeat_cnt = 0
			}
		} else {
			//build puzzle
			puzzle = model.Initialize_Puzzle()
			model.Randomize_Puzzle(puzzle, *swap_cnt)
			model.Gamify_Puzzle(puzzle, *blank_cnt)
			fmt.Printf("Created puzzle with %d swaps and %d blanks:\n", *swap_cnt, *blank_cnt)
			model.Print_Puzzle(puzzle)
		}

		//solve puzzle
		solve_start_time := time.Now()
		model.Solve_Puzzle(puzzle)
		solve_elapsed_time := time.Since(solve_start_time)
		fmt.Printf("Solved puzzle (%dms) with %d guesses:\n", solve_elapsed_time.Milliseconds(), model.Get_Solve_Stats())
		model.Print_Puzzle(puzzle)

		//reset
		model.Reset_Solve_Stats()
	}
}
