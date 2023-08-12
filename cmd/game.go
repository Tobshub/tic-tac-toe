package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tobshub/tic-tac-toe/internals"
)

var BOARD internals.Board

var (
	HasWon     = false
	GameWinner internals.XorO
)

const (
	large_font_size = 35
	small_font_size = 20
)

func InitGame(textures [2]*rl.Texture2D) {
	HasWon = false
	BOARD.Init(SCREEN_WIDTH, SCREEN_HEIGHT, textures)
}

func DrawGame() {
	BOARD.Draw()

	if HasWon {
		if GameWinner == internals.X {
			text := "Player 1 wins!"
			rl.DrawText(text, SCREEN_WIDTH/2-rl.MeasureText(text, large_font_size)/2, SCREEN_HEIGHT/2-large_font_size/2, large_font_size, rl.Red)
		} else if GameWinner == internals.O {
			text := "Player 2 wins!"
			rl.DrawText(text, SCREEN_WIDTH/2-rl.MeasureText(text, large_font_size)/2, SCREEN_HEIGHT/2-large_font_size/2, large_font_size, rl.Blue)
		}
		instruction := "Press space to restart"
		rl.DrawText(instruction, SCREEN_WIDTH/2-rl.MeasureText(instruction, small_font_size)/2, SCREEN_HEIGHT/2+large_font_size+small_font_size/2, small_font_size, rl.Gray)
	}
}

func UpdateGame(textures [2]*rl.Texture2D) {
	if !HasWon {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			has_won, winner := BOARD.Update(rl.GetMouseX(), rl.GetMouseY())

			HasWon = has_won
			if has_won {
				GameWinner = winner
			}
		}
	} else {
		if rl.IsKeyPressed(rl.KeySpace) {
			InitGame(textures)
		}
	}
}
