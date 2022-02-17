package main

import "fmt"

func Contact(numBalls int32, ball Circle, balls []Circle, numInfected *int) {
	// Update the infected population
	for i := 0; i < int(numBalls); i++ {
		for j := 1; j < int(numBalls); j++ {
			if (balls[i].centerX+int32(ball.radius) > balls[j].centerX) && (balls[i].centerX-int32(ball.radius) < balls[j].centerX) && balls[i].alive && balls[j].alive {
				if balls[i].centerY+int32(ball.radius) > balls[j].centerY && balls[i].centerY-int32(ball.radius) < balls[j].centerY {
					if balls[i].infected || balls[j].infected {
						if balls[i].infected && !balls[j].infected {
							balls[j].infected = true
							balls[j].sickTime = 0
							*numInfected++
							fmt.Println(*numInfected, " people are infected at the moment!")
						} else if !balls[i].infected && balls[j].infected {
							balls[i].infected = true
							balls[i].sickTime = 0
							*numInfected++
							fmt.Println(*numInfected, " people are infected at the moment!")
						}
					}
				}
			}
		}
	}
}
