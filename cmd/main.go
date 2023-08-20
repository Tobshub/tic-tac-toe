package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tobshub/tic-tac-toe/resources"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 640
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Tic Tac Toe")

	rl.SetTargetFPS(30)

	X_PNG := rl.LoadImageFromMemory(".png", resources.X, 5998)
	O_PNG := rl.LoadImageFromMemory(".png", resources.O, 5913)

	XTex := rl.LoadTextureFromImage(X_PNG)
	OTex := rl.LoadTextureFromImage(O_PNG)

	textures := [2]*rl.Texture2D{&XTex, &OTex}
	InitGame(textures)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		DrawGame()
		UpdateGame(textures)

		rl.EndDrawing()
	}

	rl.UnloadTexture(XTex)
	rl.UnloadTexture(OTex)

	rl.CloseWindow()
}
