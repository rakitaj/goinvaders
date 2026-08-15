package main

import "slices"

type Direction int

const (
	Up    Direction = iota // 0
	Right                  // 1
	Down                   // 2
	Left                   // 3
)

type Position struct {
	X int
	Y int
}

func NewPosition(X int, Y int) Position {
	return Position{X: X, Y: Y}
}

type GameConfig struct {
	XSize     int
	YSize     int
	MaxLength int
}

type GameState struct {
	Config            GameConfig
	Head              Position
	Length            int
	PreviousPositions []Position
	Enemies           []Position
	CurrentDirection  Direction
}

// tick: Runs the game state one iteration, moving the dragon one space
// in the current direction.
// Returns a bool indicating if the player is alive after the tick.
func (state *GameState) tick() bool {
	pos := state.Head
	switch state.CurrentDirection {
	case Up:
		pos.Y += 1
	case Right:
		pos.X += 1
	case Down:
		pos.Y -= 1
	case Left:
		pos.X -= 1
	}

	if pos.X < 0 || pos.X >= state.Config.XSize {
		return false
	}
	if pos.Y < 0 || pos.Y >= state.Config.YSize {
		return false
	}
	if slices.Contains(state.PreviousPositions, pos) {
		return false
	}

	// All the death conditions are done, time to live!
	return true
}
