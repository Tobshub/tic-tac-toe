package internals

var (
	AI_X_ON bool = false
	AI_O_ON bool = false
)

const MAX_SIMUL_DEPTH = 6

func (b *Board) filterEmptyCells() [][]int {
	empty_cells := [][]int{}
	for r := 0; r < board_r_and_c; r++ {
		for c := 0; c < board_r_and_c; c++ {
			if !b.Cells[r][c].isFilled() {
				empty_cells = append(empty_cells, []int{r, c})
			}
		}
	}
	return empty_cells
}

func (b *Board) bestEmptyCell() []int {
	simul_board := b.Copy()
	min_max := MinMax(&simul_board, b.Turn, MAX_SIMUL_DEPTH)
	return []int{min_max[1], min_max[2]}
}

func (b *Board) MakeBestMove() (bool, CellValue) {
	empty_cell_idx := b.bestEmptyCell()
	b.Cells[empty_cell_idx[0]][empty_cell_idx[1]].forceMove(b.Turn)

	has_won, winner := b.CheckWinner()
	b.nextTurn()

	return has_won, winner
}
