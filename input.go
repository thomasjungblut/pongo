package main

import (
	"github.com/nsf/termbox-go"
)

type DirectionInput struct {
	directionChannel chan Direction
}

func (keyInput *DirectionInput) Run() {
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			if input.ProcessEvent(&event) {
				break
			}
		}
	}
}

func (keyInput *DirectionInput) ProcessEvent(event *termbox.Event) bool {

	if event.Ch == 'q' || event.Key == termbox.KeyCtrlC {
		env.running = false
		return true
	}

	switch event.Key {
	case termbox.KeyArrowUp:
		input.directionChannel <- DIR_UP
		break
	case termbox.KeyArrowDown:
		input.directionChannel <- DIR_DOWN
		break
	}

	return false
}
