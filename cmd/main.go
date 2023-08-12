package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 640
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Tic Tac Toe")

	rl.SetTargetFPS(30)
	XTex := rl.LoadTexture("resources/x.png")
	OTex := rl.LoadTexture("resources/o.png")

	textures := [2]*rl.Texture2D{&XTex, &OTex}
	InitGame(textures)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		DrawGame()
		UpdateGame(textures)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
