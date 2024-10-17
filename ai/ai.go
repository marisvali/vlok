package ai

import (
	. "github.com/marisvali/vlok/gamelib"
	. "github.com/marisvali/vlok/world"
	_ "image/png"
)

type AI struct {
	frameIdx          Int
	lastRandomMoveIdx Int
}

func (a *AI) Step(w *World) (input PlayerInput) {
	return
}
