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

var ToggleAIasXButton internals.CheckBox = internals.CheckBox{
	Label:     "Toggle X AI",
	IsChecked: false,
	Color:     rl.Red,
	CheckedAction: func() {
		internals.AI_X_ON = true
		if !internals.AI_O_ON {
			internals.AI_TURN = internals.X
		}
	},
	UncheckedAction: func() {
		internals.AI_X_ON = false
		if internals.AI_TURN == internals.X {
			internals.AI_TURN = internals.O
		}
	},
}

var ToggleAIasOButton internals.CheckBox = internals.CheckBox{
	Label:     "Toggle O AI",
	IsChecked: false,
	Color:     rl.Blue,
	CheckedAction: func() {
		internals.AI_O_ON = true
		if !internals.AI_X_ON {
			internals.AI_TURN = internals.O
		}
	},
	UncheckedAction: func() {
		internals.AI_O_ON = false
		if internals.AI_TURN == internals.O {
			internals.AI_TURN = internals.X
		}
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

	ToggleAIasXButton.Size = (SCREEN_WIDTH - int32(board_right)) / 4
	ToggleAIasXButton.X = int32(board_right) + (SCREEN_WIDTH-int32(board_right))/2
	ToggleAIasXButton.Y = SCREEN_HEIGHT/2 - ToggleAIasXButton.Size

	ToggleAIasOButton.Size = (SCREEN_WIDTH - int32(board_right)) / 4
	ToggleAIasOButton.X = int32(board_right) + (SCREEN_WIDTH-int32(board_right))/2
	ToggleAIasOButton.Y = SCREEN_HEIGHT/2 + ToggleAIasXButton.Size
}

func DrawGame() {
	BOARD.Draw()
	ToggleAIasXButton.Draw()
	ToggleAIasOButton.Draw()

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
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		ToggleAIasXButton.Update(rl.GetMouseX(), rl.GetMouseY())
		ToggleAIasOButton.Update(rl.GetMouseX(), rl.GetMouseY())
	}

	if !HasWon && !IsDraw {
		if internals.AI_X_ON && internals.AI_O_ON {
			has_won, winner := BOARD.MakeBestMove()
			UpdateGameStatus(has_won, winner)
		} else if internals.AI_O_ON || internals.AI_X_ON {
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
