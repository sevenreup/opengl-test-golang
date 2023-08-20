package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	position rl.Vector2
	speed    float32
	canJump  bool
}

func main() {
var	screenWith int32 = 800
	var screenHeight int32 = 450

    
	rl.InitWindow(screenWith, screenHeight, "raylib [core] example - basic window")
    
    player := Player{rl.NewVector2(400, 280), 5, false}
    camera := rl.NewCamera2D(rl.NewVector2(float32(screenWith / 2.0), float32(screenHeight) / 2.0), player.position, 0, 1)

	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
        rl.BeginMode2D(camera);
        rl.EndMode2D();
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
