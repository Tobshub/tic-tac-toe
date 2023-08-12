package internals

import "math/rand"

func (b *Board) MakeMove() (bool, XorO) {
	empty_cells_idx := b.FilterEmptyCells() // []{r,c}

	empty_cells_len := len(empty_cells_idx)

	if empty_cells_len > 0 {
		// pick random empty spot
		cell := empty_cells_idx[rand.Intn(empty_cells_len)]
		r := cell[0]
		c := cell[1]
		b.Cells[r][c].AI_Update(b.Turn)
	}

	has_won, winner := b.CheckWinner()
	b.NextTurn()

	return has_won, winner
}
