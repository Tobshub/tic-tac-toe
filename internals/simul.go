package internals

import (
	"math"
)

func (b *Board) simulCell(cell_r, cell_c int) {
	b.Cells[cell_r][cell_c].forceMove(b.Turn)
	b.nextTurn()
}

// revert the effects of SimulCell()
func (b *Board) emptyCell(cell_r, cell_c int) {
	b.Cells[cell_r][cell_c].Value = Empty
	b.prevTurn()
}

// return {best_score, best_row, best_col}
func MinMax(b *Board, perspective CellValue, depth int) []int {
	game_over, _ := b.IsGameOver()

	var best_score int
	if b.Turn == perspective {
		best_score = math.MinInt
	} else {
		best_score = math.MaxInt
	}

	best_row := -1
	best_col := -1

	if depth <= 0 || game_over {
		best_score = EvaluateBoard(b, perspective)
	} else {
		moves := b.filterEmptyCells()

		if b.Turn == perspective {
			for _, move := range moves {
				b.simulCell(move[0], move[1])
				score := MinMax(b, perspective, depth-1)[0]
				if score > best_score {
					best_score = score
					best_row = move[0]
					best_col = move[1]
				}
				b.emptyCell(move[0], move[1])
			}
		} else if b.Turn != perspective {
			for _, move := range moves {
				b.simulCell(move[0], move[1])
				score := MinMax(b, perspective, depth-1)[0]
				if score < best_score {
					best_score = score
					best_row = move[0]
					best_col = move[1]
				}
				b.emptyCell(move[0], move[1])
			}
		}
	}

	return []int{best_score, best_row, best_col}
}

func EvaluateBoard(b *Board, perspective CellValue) int {
	score := 0

	for idx := 0; idx < board_r_and_c; idx++ {
		score += EvaluateRow(b, idx, perspective)
		score += EvaluateCol(b, idx, perspective)
	}

	score += EvaluateDiag1(b, perspective)
	score += EvaluateDiag2(b, perspective)

	return score
}

func EvaluateDiag1(b *Board, perspective CellValue) int {
	score := 0

	for idx := 0; idx < board_r_and_c; idx++ {
		should_continue, new_score := EvaluateCell(score, &b.Cells[idx][idx], idx, perspective)
		if !should_continue {
			return new_score
		} else {
			score = new_score
		}
	}

	return score
}

func EvaluateDiag2(b *Board, perspective CellValue) int {
	score := 0

	for row := 0; row < board_r_and_c; row++ {
		col := board_r_and_c - 1 - row
		should_continue, new_score := EvaluateCell(score, &b.Cells[row][col], row, perspective)
		if !should_continue {
			return new_score
		} else {
			score = new_score
		}
	}

	return score
}

func EvaluateRow(b *Board, row int, perspective CellValue) int {
	score := 0

	for col := 0; col < board_r_and_c; col++ {
		should_continue, new_score := EvaluateCell(score, &b.Cells[row][col], col, perspective)
		if !should_continue {
			return new_score
		} else {
			score = new_score
		}
	}

	return score
}

func EvaluateCol(b *Board, col int, perspective CellValue) int {
	score := 0

	for row := 0; row < board_r_and_c; row++ {
		should_continue, new_score := EvaluateCell(score, &b.Cells[row][col], row, perspective)
		if !should_continue {
			return new_score
		} else {
			score = new_score
		}
	}

	return score
}

// returns (should_continue, score)
func EvaluateCell(score int, cell *Cell, pow int, perspective CellValue) (bool, int) {
	if cell.Value == perspective {
		if score > 0 {
			return true, score * 10
		} else if score < 0 {
			return false, 0
		} else {
			return true, 1
		}
	} else if cell.Value != Empty {
		if score < 0 {
			return true, score * 10
		} else if score > 0 {
			return false, 0
		} else {
			return true, -1
		}
	}
	return true, score
}
