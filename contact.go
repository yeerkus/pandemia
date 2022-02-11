package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Contact(numInfected *int, numBalls int32, ball Circle, balls []Circle) {
	// Update the infected population
	for i := 0; i < int(numBalls); i++ {
		for j := 0; j < int(numBalls); j++ {
			if balls[i].centerX+int32(ball.radius) > balls[j].centerX && balls[i].centerX-int32(ball.radius) < balls[j].centerX {
				if balls[i].centerY+int32(ball.radius) > balls[j].centerY && balls[i].centerY-int32(ball.radius) < balls[j].centerY {
					if i != j {
						if balls[i].infected || balls[j].infected {
							if balls[i].infected && !balls[j].infected {
								balls[j].infected = true
								balls[j].color = rl.Red
								*numInfected++
								fmt.Println(*numInfected, " people are infected!")
							} else if !balls[i].infected && balls[j].infected {
								balls[i].infected = true
								balls[i].color = rl.Red
								*numInfected++
								fmt.Println(*numInfected, " people are infected!")
							}
						}
					}
				}
			}
		}
	}
}
