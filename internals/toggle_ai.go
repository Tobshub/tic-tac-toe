package internals

import rl "github.com/gen2brain/raylib-go/raylib"

type CheckBox struct {
	Size, X, Y int32
	IsChecked  bool
	Color      rl.Color
	Label      string

	CheckedAction, UncheckedAction func()
}

func (c *CheckBox) Draw() {
	font_size := int32(16)
	line_thickness := int32(4)
	rl.DrawText(c.Label, c.X+c.Size/2-rl.MeasureText(c.Label, font_size)/2, c.Y-font_size*2, font_size, rl.Black)
	rl.DrawRectangleLinesEx(
		rl.NewRectangle(float32(c.X),
			float32(c.Y), float32(c.Size),
			float32(c.Size)),
		float32(line_thickness),
		rl.Black,
	)
	if c.IsChecked {
		size_wo_thickness := c.Size - line_thickness*2
		rl.DrawRectangle(c.X+line_thickness, c.Y+line_thickness, size_wo_thickness, size_wo_thickness, c.Color)
	}
}

func (c *CheckBox) Update(mouse_x, mouse_y int32) {
	if mouse_x > c.X && mouse_x < c.X+c.Size && mouse_y > c.Y && mouse_y < c.Y+c.Size {
		if c.IsChecked {
			c.UncheckedAction()
			c.IsChecked = false
		} else {
			c.CheckedAction()
			c.IsChecked = true
		}
	}
}
