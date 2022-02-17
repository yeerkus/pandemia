package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Population(ball Circle, numBalls int32, balls []Circle, numInfected *int) {
	for i := 0; i < int(numBalls); i++ {
		// Heal the ball after 5 seconds of sickness
		if balls[i].sickTime > 300 && balls[i].alive {
			balls[i].infected = false
			balls[i].sickTime = 0
			balls[i].color = rl.Blue
			*numInfected--
		}
		if balls[i].infected {
			balls[i].color = rl.Red
			balls[i].sickTime++
		}
		// Draw the balls on the screen
		rl.DrawCircle(balls[i].centerX, balls[i].centerY, balls[i].radius, balls[i].color)

		speedX := rl.GetRandomValue(-10, 10)
		speedY := rl.GetRandomValue(-10, 10)
		// Make the balls move randomly on the screen, iff they'll not go offscreen and they're alive
		if balls[i].alive {
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
		}
	}
}
