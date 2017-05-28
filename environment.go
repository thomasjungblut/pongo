package main

type Direction int

const (
	DIR_UP   Direction = iota
	DIR_DOWN Direction = iota
)

type Position struct {
	x int
	y int
}

type Environment struct {
	running bool
	ball    *Ball
	human   *HumanPlayer
	cpu     *CPUPlayer
	xMax    int
	yMax    int
}
