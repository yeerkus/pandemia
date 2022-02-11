package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Population(ball Circle, numBalls int32, balls []Circle, ballsOnScreen *int) {
	for i := 0; i < int(numBalls); i++ {
		rl.DrawCircle(balls[i].centerX, balls[i].centerY, balls[i].radius, balls[i].color)
		speedX := rl.GetRandomValue(-10, 10)
		speedY := rl.GetRandomValue(-10, 10)
		// Make the balls move randomly on the screen, iff they'll not go offscreen
		if balls[i].centerX+speedX > int32(WIDTH) || balls[i].centerX+speedX < 0 {
			balls[i].centerX += -speedX
		} else {
			balls[i].centerX += speedX
		}
		if balls[i].centerY+speedY > int32(HEIGHT) || balls[i].centerY+speedY < 0 {
			balls[i].centerY += -speedY
		} else {
			balls[i].centerY += speedY
		}
		// Count the number of balls of the screen
		if int(balls[i].centerX) <= WIDTH && balls[i].centerY <= int32(HEIGHT) {
			if balls[i].centerX >= 0 && balls[i].centerY >= 0 {
				*ballsOnScreen++
			}
		}
	}
}
