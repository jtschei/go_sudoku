package sudoku

import "testing"

func TestInitialize_Puzzle(t *testing.T) {
	const_puzzle := Puzzle{
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
	puzzle := Initialize_Puzzle()
	if !(const_puzzle == *puzzle) {
		t.Fail()
	}

}
func TestValid_Puzzle_valid_solved(t *testing.T) {
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
	result, msg := Valid_Puzzle(&puzzle)
	if !(result == true && msg == "") {
		t.Fail()
	}
}

func TestValid_Puzzle_valid_unsolved(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 7, 2, 5, 8, 3, 6, 9},
		{2, 5, 8, 3, 6, 9, 4, 7, 1},
		{3, 6, 9, 4, 7, 1, 5, 8, 0},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 5, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	result, msg := Valid_Puzzle(&puzzle)
	if !(result == true && msg == "") {
		t.Fail()
	}
}

func TestValid_Puzzle_invalid_row_solved(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 8, 2, 5, 8, 3, 6, 9},
		{2, 5, 7, 3, 6, 9, 4, 7, 1},
		{3, 6, 9, 4, 7, 1, 5, 8, 2},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 5, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	result, msg := Valid_Puzzle(&puzzle)
	if !(result == false && msg == "Invalid row: 0") {
		t.Fail()
	}
}

func TestValid_Puzzle_invalid_row_unsolved(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 8, 2, 5, 8, 3, 6, 0},
		{2, 5, 7, 3, 6, 9, 4, 7, 1},
		{3, 6, 9, 4, 7, 1, 5, 8, 2},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 5, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	result, msg := Valid_Puzzle(&puzzle)
	if !(result == false && msg == "Invalid row: 0") {
		t.Fail()
	}
}

func TestValid_Puzzle_invalid_column_solved(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 8, 2, 5, 7, 3, 6, 9},
		{2, 5, 8, 3, 6, 9, 4, 7, 1},
		{3, 6, 9, 4, 7, 1, 5, 8, 2},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 5, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	result, msg := Valid_Puzzle(&puzzle)
	if !(result == false && msg == "Invalid column: 2") {
		t.Fail()
	}
}

func TestValid_Puzzle_invalid_column_unsolved(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 8, 2, 5, 7, 3, 6, 9},
		{2, 5, 8, 3, 6, 9, 4, 7, 1},
		{3, 6, 9, 4, 7, 1, 5, 8, 2},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 0, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	result, msg := Valid_Puzzle(&puzzle)
	if !(result == false && msg == "Invalid column: 2") {
		t.Fail()
	}
}

func TestValid_Puzzle_invalid_box_solved(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 9, 7, 2, 5, 8, 3, 6},
		{2, 5, 1, 8, 3, 6, 9, 4, 7},
		{3, 6, 2, 9, 4, 7, 1, 5, 8},
		{4, 7, 3, 1, 5, 8, 2, 6, 9},
		{5, 8, 4, 2, 6, 9, 3, 7, 1},
		{6, 9, 5, 3, 7, 1, 4, 8, 2},
		{7, 1, 6, 4, 8, 2, 5, 9, 3},
		{8, 2, 7, 5, 9, 3, 6, 1, 4},
		{9, 3, 8, 6, 1, 4, 7, 2, 5},
	}
	result, msg := Valid_Puzzle(&puzzle)
	if !(result == false && msg == "Invalid box: (row=0,col=0)") {
		t.Fail()
	}
}

func TestValid_Puzzle_invalid_box_unsolved(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 9, 7, 2, 5, 8, 3, 6},
		{2, 5, 1, 8, 3, 6, 9, 4, 7},
		{3, 0, 2, 9, 4, 7, 1, 5, 8},
		{4, 7, 3, 1, 5, 8, 2, 6, 9},
		{5, 8, 4, 2, 6, 9, 3, 7, 1},
		{6, 9, 5, 3, 7, 1, 4, 8, 2},
		{7, 1, 6, 4, 8, 2, 5, 9, 3},
		{8, 2, 7, 5, 9, 3, 6, 1, 4},
		{9, 3, 8, 6, 1, 4, 7, 2, 5},
	}
	result, msg := Valid_Puzzle(&puzzle)
	if !(result == false && msg == "Invalid box: (row=0,col=0)") {
		t.Fail()
	}
}

func TestSolved_Puzzle_valid(t *testing.T) {
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
	result := Solved_Puzzle(&puzzle)
	if result == false {
		t.Fail()
	}
}

func TestSolved_Puzzle_missing(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 7, 2, 5, 8, 3, 6, 9},
		{2, 5, 8, 3, 6, 9, 4, 7, 0},
		{3, 6, 9, 4, 7, 1, 5, 8, 2},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 5, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	result := Solved_Puzzle(&puzzle)
	if result == true {
		t.Fail()
	}
}

func TestSolved_Puzzle_invalid_row(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 8, 2, 5, 8, 3, 6, 9},
		{2, 5, 7, 3, 6, 9, 4, 7, 1},
		{3, 6, 9, 4, 7, 1, 5, 8, 2},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 5, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	result := Solved_Puzzle(&puzzle)
	if result == true {
		t.Fail()
	}
}

func TestSolved_Puzzle_invalid_column(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 8, 2, 5, 7, 3, 6, 9},
		{2, 5, 8, 3, 6, 9, 4, 7, 1},
		{3, 6, 9, 4, 7, 1, 5, 8, 2},
		{4, 7, 1, 5, 8, 2, 6, 9, 3},
		{5, 8, 2, 6, 9, 3, 7, 1, 4},
		{6, 9, 3, 7, 1, 4, 8, 2, 5},
		{7, 1, 4, 8, 2, 5, 9, 3, 6},
		{8, 2, 5, 9, 3, 6, 1, 4, 7},
		{9, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	result := Solved_Puzzle(&puzzle)
	if result == true {
		t.Fail()
	}
}

func TestSolved_Puzzle_invalid_box(t *testing.T) {
	puzzle := Puzzle{
		{1, 4, 9, 7, 2, 5, 8, 3, 6},
		{2, 5, 1, 8, 3, 6, 9, 4, 7},
		{3, 6, 2, 9, 4, 7, 1, 5, 8},
		{4, 7, 3, 1, 5, 8, 2, 6, 9},
		{5, 8, 4, 2, 6, 9, 3, 7, 1},
		{6, 9, 5, 3, 7, 1, 4, 8, 2},
		{7, 1, 6, 4, 8, 2, 5, 9, 3},
		{8, 2, 7, 5, 9, 3, 6, 1, 4},
		{9, 3, 8, 6, 1, 4, 7, 2, 5},
	}
	result := Solved_Puzzle(&puzzle)
	if result == true {
		t.Fail()
	}
}

func TestGamify_Puzzle(t *testing.T) {
	puzzle := Initialize_Puzzle()
	blank_cnt := 10
	blank_fnd := 0
	Gamify_Puzzle(puzzle, blank_cnt)
	for rowi := 0; rowi < 9; rowi++ {
		for coli := 0; coli < 9; coli++ {
			if puzzle[rowi][coli] == 0 {
				blank_fnd++
			}
		}
	}
	if !(blank_cnt == blank_fnd) {
		t.Fail()
	}
}

func TestSolve_Puzzle1(t *testing.T) {
	puzzle := Puzzle{
		{0, 0, 0, 0, 5, 0, 0, 9, 6},
		{0, 5, 4, 0, 9, 0, 0, 0, 0},
		{0, 9, 6, 0, 0, 1, 5, 0, 2},
		{7, 0, 0, 0, 4, 2, 9, 0, 3},
		{0, 4, 0, 0, 6, 0, 8, 0, 7},
		{9, 6, 3, 0, 0, 0, 0, 0, 5},
		{8, 0, 0, 0, 2, 0, 6, 3, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{6, 0, 0, 1, 0, 0, 2, 5, 4},
	}
	solved := Solve_Puzzle(&puzzle)
	if !solved {
		t.Fail()
	}
}

func TestSolve_Puzzle2(t *testing.T) {
	puzzle := Puzzle{
		{0, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 6, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{6, 0, 0, 0, 0, 2, 0, 8, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 0},
		{0, 8, 0, 0, 0, 6, 0, 2, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	solved := Solve_Puzzle(&puzzle)
	if !solved {
		t.Fail()
	}
}

func TestSolve_Puzzle3(t *testing.T) {
	puzzle := Puzzle{
		{0, 0, 0, 0, 0, 0, 0, 6, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0},
		{6, 0, 1, 0, 0, 8, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 9, 0, 0},
		{4, 0, 0, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 8, 0, 0, 5, 0},
	}
	valid, _ := Valid_Puzzle(&puzzle)
	if !valid {
		t.Fail()
	}
	solved := Solve_Puzzle(&puzzle)
	if !solved {
		t.Fail()
	}
}

func TestSolve_Puzzle4(t *testing.T) {
	puzzle := Puzzle{
		{0, 0, 0, 1, 0, 0, 0, 6, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 2},
		{0, 8, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 2, 0},
		{0, 0, 0, 6, 0, 0, 0, 1, 3},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{4, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 5, 0, 0, 0, 0, 0, 0, 0},
	}
	valid, _ := Valid_Puzzle(&puzzle)
	if !valid {
		t.Fail()
	}
	solved := Solve_Puzzle(&puzzle)
	if !solved {
		t.Fail()
	}
}

func TestSolve_Puzzle5(t *testing.T) {
	puzzle := Puzzle{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{5, 3, 0, 0, 1, 0, 0, 0, 0},
		{4, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 7, 6, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 3, 0, 0, 0},
		{2, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 9, 0, 0, 0, 2},
	}
	valid, _ := Valid_Puzzle(&puzzle)
	if !valid {
		t.Fail()
	}
	solved := Solve_Puzzle(&puzzle)
	if !solved {
		t.Fail()
	}
}
