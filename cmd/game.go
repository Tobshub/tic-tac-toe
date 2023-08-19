package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tobshub/tic-tac-toe/internals"
)

var BOARD internals.Board

var (
	HasWon     = false
	IsDraw     = false
	GameWinner internals.CellValue
)

var ToggleAIButton internals.CheckBox = internals.CheckBox{
	Label:     "Toggle AI",
	IsChecked: false,
	Color:     rl.Blue,
	CheckedAction: func() {
		internals.AI_ON = true
	},
	UncheckedAction: func() {
		internals.AI_ON = false
	},
}

const (
	large_font_size = 35
	small_font_size = 20
)

func InitGame(textures [2]*rl.Texture2D) {
	HasWon = false
	IsDraw = false
	BOARD.Init(SCREEN_WIDTH, SCREEN_HEIGHT, textures)

	board_right := BOARD.X + BOARD.Size
	ToggleAIButton.Size = (SCREEN_WIDTH - int32(board_right)) / 4
	ToggleAIButton.X = int32(board_right) + (SCREEN_WIDTH-int32(board_right))/2
	ToggleAIButton.Y = SCREEN_HEIGHT/2 - ToggleAIButton.Size/2
}

func DrawGame() {
	BOARD.Draw()
	ToggleAIButton.Draw()

	if HasWon || IsDraw {
		var text string
		if IsDraw {
			text = "It's a draw!"
			rl.DrawText(text, SCREEN_WIDTH/2-rl.MeasureText(text, large_font_size)/2, SCREEN_HEIGHT/2-large_font_size/2, large_font_size, rl.Gray)
		} else if GameWinner == internals.X {
			text = "Player 1 wins!"
			rl.DrawText(text, SCREEN_WIDTH/2-rl.MeasureText(text, large_font_size)/2, SCREEN_HEIGHT/2-large_font_size/2, large_font_size, rl.Red)
		} else if GameWinner == internals.O {
			text = "Player 2 wins!"
			rl.DrawText(text, SCREEN_WIDTH/2-rl.MeasureText(text, large_font_size)/2, SCREEN_HEIGHT/2-large_font_size/2, large_font_size, rl.Blue)
		}
		instruction := "Press space to restart"
		rl.DrawText(instruction, SCREEN_WIDTH/2-rl.MeasureText(instruction, small_font_size)/2, SCREEN_HEIGHT/2+large_font_size+small_font_size/2, small_font_size, rl.Gray)
	}
}

func UpdateGameStatus(has_won bool, winner internals.CellValue) {
	if has_won {
		HasWon = has_won
		GameWinner = winner
	} else {
		is_draw := BOARD.CheckDrawState()
		if is_draw {
			IsDraw = true
		}
	}
}

func UpdateGame(textures [2]*rl.Texture2D) {
	if !HasWon && !IsDraw {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			ToggleAIButton.Update(rl.GetMouseX(), rl.GetMouseY())
		}
		if internals.AI_ON {
			if BOARD.Turn != internals.AI_TURN && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				has_won, winner := BOARD.Update(rl.GetMouseX(), rl.GetMouseY())
				UpdateGameStatus(has_won, winner)
			} else if BOARD.Turn == internals.AI_TURN {
				has_won, winner := BOARD.MakeBestMove()
				UpdateGameStatus(has_won, winner)
			}
		} else {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				has_won, winner := BOARD.Update(rl.GetMouseX(), rl.GetMouseY())
				UpdateGameStatus(has_won, winner)
			}
		}
	} else {
		if rl.IsKeyPressed(rl.KeySpace) {
			InitGame(textures)
		}
	}
}
