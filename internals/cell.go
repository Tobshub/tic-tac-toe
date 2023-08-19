package internals

import rl "github.com/gen2brain/raylib-go/raylib"

type CellValue int

const (
	X     CellValue = -1
	Empty CellValue = 0
	O     CellValue = 1
)

type Cell struct {
	Value CellValue

	X, Y float32
}

func (c *Cell) isFilled() bool {
	return c.Value != Empty
}

func (c *Cell) Draw(cell_size float32, textures [2]*rl.Texture2D) {
	x := c.X
	y := c.Y

	rl.DrawRectangleLinesEx(rl.NewRectangle(x, y, cell_size, cell_size), 2, rl.Black)
	XTex := textures[0]
	OTex := textures[1]

	if c.isFilled() {
		switch c.Value {
		case X:
			rl.DrawTextureRec(
				*XTex,
				rl.NewRectangle(0, 0, float32(XTex.Width), float32(XTex.Height)),
				rl.NewVector2(x+cell_size/2-float32(XTex.Width)/2, y+cell_size/2-float32(XTex.Height)/2),
				rl.White,
			)
		case O:
			rl.DrawTextureRec(
				*OTex,
				rl.NewRectangle(0, 0, float32(OTex.Width), float32(OTex.Height)),
				rl.NewVector2(x+cell_size/2-float32(OTex.Width)/2, y+cell_size/2-float32(OTex.Height)/2),
				rl.White,
			)
		}
	}
}

func (c *Cell) Update(turn CellValue, mouse_x, mouse_y, cell_size float32) bool {
	if c.isFilled() {
		return false
	}
	if mouse_x >= c.X && mouse_x <= c.X+cell_size && mouse_y >= c.Y && mouse_y <= c.Y+cell_size {
		c.Value = turn
		return true
	} else {
		return false
	}
}

func (c *Cell) forceMove(turn CellValue) {
	if !c.isFilled() {
		c.Value = turn
	}
}

func (c *Cell) copy() Cell {
	return Cell{
		Value: c.Value,
		X:     c.X,
		Y:     c.Y,
	}
}
