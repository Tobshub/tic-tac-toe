package internals

import (
	"math"
)

func (b *Board) SimulCell(cell_r, cell_c int) {
	b.Cells[cell_r][cell_c].ForceMove(b.Turn)
	b.NextTurn()
}

func (b *Board) EmptyCell(cell_r, cell_c int) {
	b.Cells[cell_r][cell_c].Value = Empty
	b.PrevTurn()
}

// return {best_score, best_row, best_col}
func MinMax(b *Board, depth int) []int {
	game_over, _ := b.IsGameOver()

	var best_score int
	if b.Turn == AI_TURN {
		best_score = math.MinInt
	} else {
		best_score = math.MaxInt
	}

	best_row := -1
	best_col := -1

	if depth <= 0 || game_over {
		best_score = EvaluateBoard(b)
	} else {
		moves := b.FilterEmptyCells()

		if b.Turn == AI_TURN {
			for _, move := range moves {
				b.SimulCell(move[0], move[1])
				score := MinMax(b, depth-1)[0]
				if score > best_score {
					best_score = score
					best_row = move[0]
					best_col = move[1]
				}
				b.EmptyCell(move[0], move[1])
			}
		} else if b.Turn != AI_TURN {
			for _, move := range moves {
				b.SimulCell(move[0], move[1])
				score := MinMax(b, depth-1)[0]
				if score < best_score {
					best_score = score
					best_row = move[0]
					best_col = move[1]
				}
				b.EmptyCell(move[0], move[1])
			}
		}
	}

	return []int{best_score, best_row, best_col}
}

func EvaluateBoard(b *Board) int {
	score := 0

	for idx := 0; idx < board_r_and_c; idx++ {
		score += EvaluateRow(b, idx)
		score += EvaluateCol(b, idx)
	}

	score += EvaluateDiag1(b)
	score += EvaluateDiag2(b)

	return score
}

func EvaluateDiag1(b *Board) int {
	score := 0

	for idx := 0; idx < board_r_and_c; idx++ {
		if b.Cells[idx][idx].Value == AI_TURN {
			if score > 0 {
				score = int(math.Pow10(idx))
			} else if score < 0 {
				return 0
			} else {
				score = 1
			}
		} else if b.Cells[idx][idx].Value != Empty {
			if score < 0 {
				score = int(math.Pow10(idx)) * -1
			} else if score > 0 {
				return 0
			} else {
				score = -1
			}
		}
	}

	return score
}

func EvaluateDiag2(b *Board) int {
	score := 0

	for row := 0; row < board_r_and_c; row++ {
		col := board_r_and_c - 1 - row
		if b.Cells[row][col].Value == AI_TURN {
			if score > 0 {
				score = int(math.Pow10(row))
			} else if score < 0 {
				return 0
			} else {
				score = 1
			}
		} else if b.Cells[row][col].Value != Empty {
			if score < 0 {
				score = int(math.Pow10(row)) * -1
			} else if score > 0 {
				return 0
			} else {
				score = -1
			}
		}
	}

	return score
}

func EvaluateRow(b *Board, row int) int {
	score := 0

	for col := 0; col < board_r_and_c; col++ {
		if b.Cells[row][col].Value == AI_TURN {
			if score > 0 {
				score = int(math.Pow10(col))
			} else if score < 0 {
				return 0
			} else {
				score = 1
			}
		} else if b.Cells[row][col].Value != Empty {
			if score < 0 {
				score = int(math.Pow10(col)) * -1
			} else if score > 0 {
				return 0
			} else {
				score = -1
			}
		}
	}

	return score
}

func EvaluateCol(b *Board, col int) int {
	score := 0

	for row := 0; row < board_r_and_c; row++ {
		if b.Cells[row][col].Value == AI_TURN {
			if score > 0 {
				score = int(math.Pow10(row))
			} else if score < 0 {
				return 0
			} else {
				score = 1
			}
		} else if b.Cells[row][col].Value != Empty {
			if score < 0 {
				score = int(math.Pow10(row)) * -1
			} else if score > 0 {
				return 0
			} else {
				score = -1
			}
		}
	}

	return score
}
