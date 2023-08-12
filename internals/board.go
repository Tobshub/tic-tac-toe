package internals

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screen_padding int32 = 200
	board_r_and_c        = 3
)

type Board struct {
	Cells                [board_r_and_c][board_r_and_c]Cell
	Size, X, Y, CellSize float32
	Turn                 XorO
	Textures             [2]*rl.Texture2D
}

func (b *Board) Init(screen_width, screen_height int32, textures [2]*rl.Texture2D) {
	b.Cells = [board_r_and_c][board_r_and_c]Cell{}
	b.Turn = X

	b.Textures = textures

	screen_min := math.Min(float64(screen_width), float64(screen_height))

	b.Size = float32(screen_min) - float32(screen_padding)

	b.X = float32(screen_width/2) - b.Size/2
	b.Y = float32(screen_height/2) - b.Size/2

	b.CellSize = b.Size / float32(board_r_and_c)

	for r := 0; r < board_r_and_c; r++ {
		for c := 0; c < board_r_and_c; c++ {
			b.Cells[r][c] = Cell{X: b.X + float32(c)*b.CellSize, Y: b.Y + float32(r)*b.CellSize, Filled: false}
		}
	}
}

func (b *Board) Draw() {
	rl.DrawRectangleLinesEx(rl.NewRectangle(b.X, b.Y, b.Size, b.Size), 4, rl.Black)
	b.DrawCells()
}

func (b *Board) DrawCells() {
	for r := 0; r < board_r_and_c; r++ {
		for c := 0; c < board_r_and_c; c++ {
			b.Cells[r][c].Draw(b.CellSize, b.Textures)
		}
	}
}

func (b *Board) Update(mouse_x, mouse_y int32) (bool, XorO) {
	for r := 0; r < board_r_and_c; r++ {
		for c := 0; c < board_r_and_c; c++ {
			done := b.Cells[r][c].Update(b.Turn, float32(mouse_x), float32(mouse_y), b.CellSize)
			if done {
				has_won, winner := b.CheckWinner()
				b.NextTurn()
				return has_won, winner
			}
		}
	}
	return false, b.Turn
}

func (b *Board) NextTurn() {
	switch b.Turn {
	case X:
		b.Turn = O
	case O:
		b.Turn = X
	}
}

func (b *Board) CheckWinner() (bool, XorO) {
	for r := 0; r < board_r_and_c; r++ {
		check_row := true
		check_col := true

		for c := 0; c < board_r_and_c; c++ {
			if b.Cells[r][c].Filled {
				if c > 0 && c < board_r_and_c-1 {
					check_row = check_row && b.Cells[r][c].Value == b.Cells[r][c-1].Value && b.Cells[r][c].Value == b.Cells[r][c+1].Value
				} else if c == 0 {
					check_row = check_row && b.Cells[r][c].Value == b.Cells[r][c+1].Value
				} else {
					check_row = check_row && b.Cells[r][c].Value == b.Cells[r][c-1].Value
				}
			} else {
				check_row = false
			}

			if b.Cells[c][r].Filled {
				if c > 0 && c < board_r_and_c-1 {
					check_col = check_col && b.Cells[c][r].Value == b.Cells[c-1][r].Value && b.Cells[c][r].Value == b.Cells[c+1][r].Value
				} else if c == 0 {
					check_col = check_col && b.Cells[c][r].Value == b.Cells[c+1][r].Value
				} else {
					check_col = check_col && b.Cells[c][r].Value == b.Cells[c-1][r].Value
				}
			} else {
				check_col = false
			}
		}

		if check_row {
			return true, b.Cells[r][0].Value
		} else if check_col {
			return true, b.Cells[0][r].Value
		}
	}

	check_diag_1 := true

	for r := 0; r < board_r_and_c; r++ {
		if b.Cells[r][r].Filled {
			if r > 0 && r < board_r_and_c-1 {
				check_diag_1 = check_diag_1 && b.Cells[r][r].Value == b.Cells[r-1][r-1].Value && b.Cells[r][r].Value == b.Cells[r+1][r+1].Value
			} else if r == 0 {
				check_diag_1 = check_diag_1 && b.Cells[r][r].Value == b.Cells[r+1][r+1].Value
			} else {
				check_diag_1 = check_diag_1 && b.Cells[r][r].Value == b.Cells[r-1][r-1].Value
			}
		} else {
			check_diag_1 = false
		}
	}

	if check_diag_1 {
		return true, b.Cells[0][0].Value
	}

	check_diag_2 := true

	for r := 0; r < board_r_and_c; r++ {
		c := board_r_and_c - r - 1
		if b.Cells[r][c].Filled {
			if r > 0 && r < board_r_and_c-1 {
				check_diag_2 = check_diag_2 && b.Cells[r][c].Value == b.Cells[r+1][c-1].Value && b.Cells[r][r].Value == b.Cells[r-1][c+1].Value
			} else if r == 0 {
				check_diag_2 = check_diag_2 && b.Cells[r][c].Value == b.Cells[r+1][c-1].Value
			} else {
				check_diag_2 = check_diag_2 && b.Cells[r][c].Value == b.Cells[r-1][c+1].Value
			}
		} else {
			check_diag_2 = false
		}
	}

	if check_diag_2 {
		return true, b.Cells[0][board_r_and_c-1].Value
	}

	return false, b.Turn
}

func (b *Board) CheckDrawState() bool {
	for r := 0; r < board_r_and_c; r++ {
		for c := 0; c < board_r_and_c; c++ {
			if !b.Cells[r][c].Filled {
				return false
			}
		}
	}

	has_won, _ := b.CheckWinner()

	if has_won {
		return false
	}

	return true
}

func (b *Board) FilterEmptyCells() [][]int {
	empty_cells := [][]int{}
	for r := 0; r < board_r_and_c; r++ {
		for c := 0; c < board_r_and_c; c++ {
			if !b.Cells[r][c].Filled {
				empty_cells = append(empty_cells, []int{r, c})
			}
		}
	}
	return empty_cells
}
