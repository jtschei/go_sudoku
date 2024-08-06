package sudoku

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

/*
Puzzle board consisting of 9x9 grid
First index is the row.
Second index is the column.
*/
type Puzzle [9][9]uint8

/*
A blank position in puzzle and its possible values
*/
type blank_option_type struct {
	blank   [2]uint8 //[row,column]
	options []uint8  //[possiblevalue,...]
}

/*
Constructs a constant, valid, and solved puzzle.
*/
func Initialize_Puzzle() *Puzzle {
	puzzle := Puzzle{
		{1, 4, 7, 2, 5, 8, 3, 6, 9},
		{2, 5, 8, 3, 6, 9, 4, 7, 1},
		{3, 6, 9, 4, 7, 1, 5, 8, 2},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 5, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	return &puzzle
}

/*
Tracks history of guesses employed solving puzzle (memoization).
*/
var guess_cnt int = 0
var failed_guess_history map[[81]uint8]bool = make(map[[81]uint8]bool)

/*
Prints a puzzle to standard out.
*/
const puzzle_print_header = "/---------|---------|---------\\"
const puzzle_print_col_sep = "|"
const puzzle_print_col_pad = " "
const puzzle_print_row_sep = "|---------|---------|---------|"
const puzzle_print_footer = "\\---------|---------|---------/"

func Print_Puzzle(puzzle *Puzzle) {
	fmt.Println(puzzle_print_header)
	for rowi := range puzzle {
		fmt.Print(puzzle_print_col_sep)
		for coli, volv := range puzzle[rowi] {
			fmt.Print(puzzle_print_col_pad, volv, puzzle_print_col_pad)
			if coli == 2 || coli == 5 {
				fmt.Print(puzzle_print_col_sep)
			}
		}
		fmt.Println(puzzle_print_col_sep)
		if rowi == 2 || rowi == 5 {
			fmt.Println(puzzle_print_row_sep)
		}
	}
	fmt.Println(puzzle_print_footer)
	fmt.Println()
}

/*
Reads a puzzle in printed format and builds internal representation.
*/
func Puzzle_From_String(puzzle_str string) *Puzzle {
	var puzzle Puzzle
	lines := strings.Split(puzzle_str, "\n")
	rowi := 0
	for _, line := range lines {
		if line == puzzle_print_header {
			continue
		} else if line == puzzle_print_row_sep {
			continue
		} else if line == puzzle_print_footer {
			continue
		} else if len(line) == 0 {
			continue
		} else {
			//remove the column separator and extra spaces
			line = strings.ReplaceAll(line, puzzle_print_col_sep+puzzle_print_col_pad, "")
			line = strings.ReplaceAll(line, puzzle_print_col_pad+puzzle_print_col_sep, "")
			line = strings.ReplaceAll(line, puzzle_print_col_pad+puzzle_print_col_pad, " ")
			//splint row into values split by space
			vals := strings.Split(line, puzzle_print_col_pad)
			//incorporate values into the puzzle's row
			for vali, val := range vals {
				if vali < 9 {
					val_int, _ := strconv.Atoi(val)
					puzzle[rowi][vali] = uint8(val_int)
				}
			}
			rowi++
		}
	}
	valid, msg := Valid_Puzzle(&puzzle)
	if !valid {
		fmt.Println("Puzzle is invalid: ", msg)
		return nil
	}
	return &puzzle
}

/*
Identifies positions in puzzle that can swapped (with another swap set).
A value is randomly chosen in first box and its position saved.
The position of the value in remaining boxes is additionaly saved.
*/
func get_swaps(puzzle *Puzzle) [9][2]uint8 {
	var swaps [9][2]uint8
	var swapi int = 0
	swaps[swapi][0] = uint8(rand.Intn(3))
	swaps[swapi][1] = uint8(rand.Intn(3))
	swapi++
	for rowi := range puzzle {
		for coli := range puzzle[rowi] {
			if rowi < 3 && coli < 3 {
				continue
			}
			if puzzle[swaps[0][0]][swaps[0][1]] == puzzle[rowi][coli] {
				swaps[swapi][0] = uint8(rowi)
				swaps[swapi][1] = uint8(coli)
				swapi++
			}
		}
	}
	return swaps
}

/*
Will randomly swap positions of values in puzzle swap_cnt times.
The result will be a solved puzzle with random (yet valid) value positioning.
*/
func Randomize_Puzzle(puzzle *Puzzle, swap_cnt int) {
	for swapi := 0; swapi < swap_cnt; swapi++ {
		swaps1 := get_swaps(puzzle)
		swaps2 := get_swaps(puzzle)
		//make sure swaps are not same
		if swaps1 == swaps2 {
			//try again
			swapi--
			continue
		}
		//perform the swaps in the puzzle
		for swapi := 0; swapi < 9; swapi++ {
			v1 := puzzle[swaps1[swapi][0]][swaps1[swapi][1]]
			v2 := puzzle[swaps2[swapi][0]][swaps2[swapi][1]]
			puzzle[swaps1[swapi][0]][swaps1[swapi][1]] = v2
			puzzle[swaps2[swapi][0]][swaps2[swapi][1]] = v1
		}
	}
}

/*
Tests a given puzzle and returns true if puzzle is properly solved.
*/
func Solved_Puzzle(puzzle *Puzzle) bool {
	var vals [9]uint8
	//test each row
	for rowi := 0; rowi < 9; rowi++ {
		vals = [len(vals)]uint8{}
		for coli := 0; coli < 9; coli++ {
			v := puzzle[rowi][coli]
			if v > 0 {
				vals[v-1]++
			}
		}
		//verify
		for _, v := range vals {
			if v == 0 || v > 1 {
				return false
			}
		}
	}
	//test each column
	for coli := 0; coli < 9; coli++ {
		vals = [len(vals)]uint8{}
		for rowi := 0; rowi < 9; rowi++ {
			v := puzzle[rowi][coli]
			if v > 0 {
				vals[v-1]++
			}
		}
		//verify
		for _, v := range vals {
			if v == 0 || v > 1 {
				return false
			}
		}
	}
	//test each box
	for browi := 0; browi < 3; browi++ {
		for bcoli := 0; bcoli < 3; bcoli++ {
			vals = [len(vals)]uint8{}
			//for each slot in box
			for rowi := 3 * browi; rowi < browi*3+3; rowi++ {
				for coli := 3 * bcoli; coli < bcoli*3+3; coli++ {
					v := puzzle[rowi][coli]
					if v > 0 {
						vals[v-1]++
					}
				}
			}
			//verify
			for _, v := range vals {
				if v == 0 || v > 1 {
					return false
				}
			}
		}
	}
	return true
}

/*
Tests a given puzzle and returns true if puzzle has no invalid non-empty values.
If an invalid value is encountered, the error message will be non-blank and describe infraction.
*/
func Valid_Puzzle(puzzle *Puzzle) (bool, string) {
	var vals [9]uint8
	//test each row
	for rowi := 0; rowi < 9; rowi++ {
		vals = [len(vals)]uint8{}
		for coli := 0; coli < 9; coli++ {
			v := puzzle[rowi][coli]
			if v > 0 {
				vals[v-1]++
			}
		}
		//verify
		for _, v := range vals {
			if v > 1 {
				return false, fmt.Sprintf("Invalid row: %d", rowi)
			}
		}
	}
	//test each column
	for coli := 0; coli < 9; coli++ {
		vals = [len(vals)]uint8{}
		for rowi := 0; rowi < 9; rowi++ {
			v := puzzle[rowi][coli]
			if v > 0 {
				vals[v-1]++
			}
		}
		//verify
		for _, v := range vals {
			if v > 1 {
				return false, fmt.Sprintf("Invalid column: %d", coli)
			}
		}
	}
	//test each box
	for browi := 0; browi < 3; browi++ {
		for bcoli := 0; bcoli < 3; bcoli++ {
			vals = [len(vals)]uint8{}
			//for each slot in box
			for rowi := 3 * browi; rowi < browi*3+3; rowi++ {
				for coli := 3 * bcoli; coli < bcoli*3+3; coli++ {
					v := puzzle[rowi][coli]
					if v > 0 {
						vals[v-1]++
					}
				}
			}
			//verify
			for _, v := range vals {
				if v > 1 {
					return false, fmt.Sprintf("Invalid box: (row=%d,col=%d)", browi, bcoli)
				}
			}
		}
	}
	return true, ""
}

/*
Randomly removes blank_cnt number of positions in the puzzle and replaces value with zero (blank).
*/
func Gamify_Puzzle(puzzle *Puzzle, blank_cnt int) {
	if blank_cnt <= 0 {
		return
	}
	if blank_cnt > 81 {
		blank_cnt = 81
	}
	for blanked := 0; blanked < blank_cnt; {
		row := uint8(rand.Intn(9))
		col := uint8(rand.Intn(9))
		if puzzle[row][col] != 0 {
			puzzle[row][col] = 0
			blanked++
		}
	}
}

/*
Returns the position of all blanks within the puzzle.
*/
func get_blanks(puzzle *Puzzle) [][2]uint8 {
	var blanks [][2]uint8
	for rowi := 0; rowi < 9; rowi++ {
		for coli := 0; coli < 9; coli++ {
			if puzzle[rowi][coli] == 0 {
				blanks = append(blanks, [2]uint8{uint8(rowi), uint8(coli)})
			}
		}
	}
	return blanks
}

/*
Returns the possible values for a blank position in the puzzle.
*/
func get_options(puzzle *Puzzle, row uint8, col uint8) []uint8 {
	vals_row := [9]uint8{}
	vals_col := [9]uint8{}
	vals_box := [9]uint8{}
	var options = make([]uint8, 0, 9)
	//var options []uint8

	//each column in the row
	for coli := 0; coli < 9; coli++ {
		v := puzzle[row][coli]
		if v != 0 {
			vals_row[v-1]++
		}
	}
	//each row in the column
	for rowi := 0; rowi < 9; rowi++ {
		v := puzzle[rowi][col]
		if v != 0 {
			vals_col[v-1]++
		}
	}

	//each row and column
	// var rowi, coli uint8
	// for rowi = 0; rowi < 9; rowi++ {
	// 	for coli = 0; coli < 9; coli++ {
	// 		v := puzzle[rowi][coli]
	// 		if rowi == row && v != 0 {
	// 			vals_row[v-1]++
	// 		}
	// 		if coli == col && v != 0 {
	// 			vals_col[v-1]++
	// 		}
	// 	}
	// }

	//each box
	browi := uint8(row / 3)
	bcoli := uint8(col / 3)
	var rowi, coli uint8
	for rowi = 3 * browi; rowi < 3*browi+3; rowi++ {
		for coli = 3 * bcoli; coli < 3*bcoli+3; coli++ {
			v := puzzle[rowi][coli]
			if v != 0 {
				vals_box[v-1]++
			}
		}
	}

	//intersection
	var vi uint8
	for vi = 0; vi < 9; vi++ {
		if vals_row[vi]|vals_col[vi]|vals_box[vi] == 0 {
			options = append(options, vi+1)
		}
	}

	return options
}

/*
Converts a puzzle into an map key.
*/
func puzzle_to_key(puzzle *Puzzle) [81]uint8 {
	var puzzle_vals [81]uint8
	val_idx := 0
	for rowi := 0; rowi < 9; rowi++ {
		for coli := 0; coli < 9; coli++ {
			puzzle_vals[val_idx] = puzzle[rowi][coli]
			val_idx++
		}
	}
	return puzzle_vals
}

/*
Sorts all the blank positions and their possible values.
The sort is by:
1) length of options (asc)
2) row position (asc)
3) column position (asc)
*/
func sort_blanks_options(blanks_options []blank_option_type) {
	sort.Slice(blanks_options, func(i, j int) bool {
		if len(blanks_options[i].options) == len(blanks_options[j].options) {
			if blanks_options[i].blank[0] == blanks_options[j].blank[0] {
				return blanks_options[i].blank[1] < blanks_options[j].blank[1]
			} else {
				return blanks_options[i].blank[0] < blanks_options[j].blank[0]
			}
		} else {
			return len(blanks_options[i].options) < len(blanks_options[j].options)
		}
	})
}

/*
Returns all the blank positions and their possible values.
The resulting slice is sorted.
If the puzzle has an unsolveable blank position and empty slice is returned.
*/
func get_blanks_options(puzzle *Puzzle) []blank_option_type {
	var blanks_options = make([]blank_option_type, 0, 81)
	//var blanks_options []blank_option_type

	//get all blanks and options
	blanks := get_blanks(puzzle)
	for _, blank := range blanks {
		options := get_options(puzzle, blank[0], blank[1])
		if options == nil {
			return []blank_option_type{}
		}
		blanks_options = append(blanks_options, blank_option_type{blank, options})
	}

	//sort
	sort_blanks_options(blanks_options)

	return blanks_options
}

/*
Copies one puzzle to another puzzle.
*/
func copy_puzzle(src *Puzzle, dst *Puzzle) {
	for rowi := 0; rowi < 9; rowi++ {
		for coli := 0; coli < 9; coli++ {
			dst[rowi][coli] = src[rowi][coli]
		}
	}
}

/*
For a puzzle with a blank with >1 options, this method will "guess" a possible option for a blank.
If the puzzle guess results in the puzzle being solved the return value will be true, otherwise false.
This method keeps track of its guesses via the guess_history map.
*/
func guess_solve(puzzle *Puzzle) bool {

	//get guesses to pursue
	blanks_options := get_blanks_options(puzzle)
	if len(blanks_options) == 0 {
		//unsolveable puzzle
		return false
	}

	//iterate through the guess options
	for _, blank_option := range blanks_options {

		//test if this option has already been guessed
		_, present := failed_guess_history[puzzle_to_key(puzzle)]
		if present {
			return false
		}

		//iterate through the possible values for particular blank position
		for _, option := range blank_option.options {

			//copy puzzle before we make the guess
			var puzzle_pre_guess Puzzle
			copy_puzzle(puzzle, &puzzle_pre_guess)

			//guess
			puzzle[blank_option.blank[0]][blank_option.blank[1]] = option

			//copy puzzle after guess is made
			var puzzle_post_guess Puzzle
			copy_puzzle(puzzle, &puzzle_post_guess)

			//test if we have already attempted this guess
			_, present := failed_guess_history[puzzle_to_key(puzzle)]
			if !present {

				//try to solve with guess
				guess_cnt++
				if Solve_Puzzle(puzzle) {
					return true
				}

				//save the guess history
				failed_guess_history[puzzle_to_key(&puzzle_post_guess)] = false
			}

			//the blank's guess with paritcular option was unsuccessful
			//restore the puzzle pre-guess
			copy_puzzle(&puzzle_pre_guess, puzzle)
		}

		//no guess for blank resulted in solved puzzle
		//save history
		failed_guess_history[puzzle_to_key(puzzle)] = false
	}

	//no more guesses for puzzle
	return false
}

/*
Identifies blanks in a puzzle that have only one valid option.
Substitutes blank with appropriate value.
Returns number of blanks substituted.
*/
func deduce_solve(puzzle *Puzzle) {
	solved_cnt := 0

	for {

		//get blanks/options
		blanks_options := get_blanks_options(puzzle)
		if len(blanks_options) == 0 {
			//no further options
			return
		}

		//deduce
		for _, blank_option := range blanks_options {
			if len(blank_option.options) == 1 {
				//incorporate single option and repeat process of getting blanks/options
				puzzle[blank_option.blank[0]][blank_option.blank[1]] = blank_option.options[0]
				solved_cnt++
				//try next
				break
			} else {
				//no more single option deductions available so return
				return
			}
		}

	}
}

/*
Will solve the given puzzle so that when true is returned, the resulting puzzle is valid.
*/
func Solve_Puzzle(puzzle *Puzzle) bool {
	deduce_solve(puzzle)
	solved := Solved_Puzzle(puzzle)
	if solved {
		return true
	} else {
		return guess_solve(puzzle)
	}
}

func Get_Solve_Stats() int {
	return guess_cnt
}

func Reset_Solve_Stats() {
	guess_cnt = 0
	failed_guess_history = make(map[[81]uint8]bool)
}
