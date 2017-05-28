package main

import (
	"time"
	"math"
	"math/rand"
)

const (
	// Yes kids, I'm tying the framerate to the game logic- KMA.
	FPS       = 30
	BallSpeed = 2.0
)

type Ball struct {
	ballPositions chan Position
	position      Position
	direction     float64 // in radians
}

func (b *Ball) RunBall(env Environment) {
	b.direction = RandomBallDirection()

	for env.running {
		start := time.Now()
		// every second we recompute the balls movement
		b.position.x += int(math.Cos(b.direction) * BallSpeed)
		b.position.y += int(math.Sin(b.direction) * BallSpeed)

		if b.position.y <= 0 || b.position.y >= (env.yMax-1) {
			b.direction = 2*math.Pi - b.direction
		}

		// reset if it was in the goal area
		if b.position.x <= 0 || b.position.x >= env.xMax {
			b.position.x = env.xMax / 2
			b.position.y = env.yMax / 2
			b.direction = RandomBallDirection()
		}

		if Intersects(b.position, env.human.position) ||
			Intersects(b.position, env.cpu.position) {
			b.direction = math.Pi - b.direction
		}

		// notify about the ball position
		b.ballPositions <- Position{x: b.position.x, y: b.position.y}

		duration := time.Now().Sub(start)
		time.Sleep((time.Duration(1000/FPS) * time.Millisecond) - duration)
	}
}

type HumanPlayer struct {
	playerInput *DirectionInput
	position    Position
}

func (p *HumanPlayer) RunPlayer(env Environment) {
	for env.running {
		dir := <-p.playerInput.directionChannel

		switch dir {
		case DIR_UP:
			env.human.position.y = Max(env.human.position.y-1, 0)
			break
		case DIR_DOWN:
			env.human.position.y = Min(env.human.position.y+1, env.yMax-PaddleHeight)
		}
	}
}

type CPUPlayer struct {
	ballPositions chan Position
	position      Position
}

func (p *CPUPlayer) RunPlayer(env Environment) {
	for env.running {
		ballPos := <-p.ballPositions

		// we are basically cheating and taking the y coordinate of the ball
		// and warp along the y axis to always reach it ;-)
		env.cpu.position.y = ballPos.y - (PaddleHeight / 2)
	}
}

func RandomBallDirection() float64 {
	offset := 0.5
	if rand.Float64() > 0.5 {
		offset += 1.0
	}
	return (offset + 0.3) * math.Pi
}

func Intersects(ball Position, paddle Position) bool {

	if ball.x >= paddle.x &&
		ball.x <= paddle.x+PaddleWidth &&
		ball.y >= paddle.y &&
		ball.y <= paddle.y+PaddleHeight {
		return true
	}

	return false
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
