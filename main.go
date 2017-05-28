package main

import (
	tb "github.com/nsf/termbox-go"
)

const (
	BlankColor   = tb.ColorBlack
	PaddleWidth  = 2  // how many pixels right from the top
	PaddleHeight = 5  // how many pixels down from the top
	PaddleOffset = 10 // how many pixels from the side to offset the paddle
)

var (
	env   *Environment
	input *DirectionInput
)

func render() {
	for env.running {

		tb.Clear(BlankColor, BlankColor)

		DrawCircle(env.ball.position.x, env.ball.position.y)
		DrawPaddle(env.human.position.x, env.human.position.y)
		DrawPaddle(env.cpu.position.x, env.cpu.position.y)

		tb.Flush()
	}
}

func DrawCircle(x int, y int) {
	tb.SetCell(x, y, 'o', tb.ColorWhite, tb.ColorBlack)
}

func DrawPaddle(x int, y int) {
	for xi := x; xi < x+PaddleWidth; xi++ {
		for yi := y; yi < y+PaddleHeight; yi++ {
			tb.SetCell(xi, yi, ' ', tb.ColorDefault, tb.ColorWhite)
		}
	}
}

func main() {
	err := tb.Init()
	if err != nil {
		panic(err)
	}

	tb.SetInputMode(tb.InputEsc)
	tb.Flush()

	x, y := tb.Size()
	input = &DirectionInput{directionChannel: make(chan Direction, 1) }
	ball := &Ball{
		position:      Position{x: x / 2, y: y / 2},
		ballPositions: make(chan Position, 1),
	}
	env = &Environment{
		running: true,
		xMax:    x,
		yMax:    y,
		ball:    ball,
		human:   &HumanPlayer{position: Position{x: PaddleOffset, y: y / 2}, playerInput: input},
		cpu:     &CPUPlayer{position: Position{x: x - PaddleOffset, y: y / 2}, ballPositions: ball.ballPositions},
	}

	go ball.RunBall(*env)
	go input.Run()
	go env.human.RunPlayer(*env)
	go env.cpu.RunPlayer(*env)

	render()

}
