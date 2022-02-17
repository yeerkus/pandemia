package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  int = 800
	HEIGHT int = 600
)

type Circle struct {
	centerX  int32
	centerY  int32
	radius   float32
	color    rl.Color
	infected bool
	sickTime int32
	alive    bool
}

func main() {
	rl.SetTargetFPS(60)
	// Create a ball object to be used to manipulate each member of the population
	ball := Circle{centerX: int32(WIDTH) / 2, centerY: int32(HEIGHT) / 2, radius: 10, color: rl.White, infected: false, sickTime: 0, alive: true}
	rl.InitWindow(int32(WIDTH), int32(HEIGHT), "Go Simulation")
	// Random value of the population
	var numBalls int32 = rl.GetRandomValue(10, 100)
	// Array to contain everyone in the population
	balls := make([]Circle, numBalls)
	// Infected ball
	balls[0] = Circle{centerX: rl.GetRandomValue(0, int32(WIDTH)), centerY: rl.GetRandomValue(0, int32(HEIGHT)), radius: ball.radius, color: rl.Red, infected: true, alive: true}
	var numInfected int = 1
	// Populate the array with Ball objects. Give each person a random position
	for i := 1; i < int(numBalls); i++ {
		balls[i] = Circle{centerX: rl.GetRandomValue(0, int32(WIDTH)), centerY: rl.GetRandomValue(0, int32(HEIGHT)), radius: ball.radius, color: ball.color, infected: false, alive: true}
	}
	// MAIN Game Loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// Draw the percentange of the population that is infected
		rl.DrawText(strconv.Itoa(int(numBalls))+"  "+strconv.Itoa(100*numInfected/int(numBalls))+"%", 5*int32(WIDTH)/12, 5*int32(HEIGHT)/12, 90, rl.Beige)
		//fmt.Println("  " + strconv.Itoa(100*ballsOnScreen/int(numBalls)) + "%")
		// Call the population function: Draw the population on the screen and do it's logic
		Population(ball, numBalls, balls, &numInfected)
		// Check for collisions and update the infected
		Contact(int32(numBalls), ball, balls, &numInfected)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
