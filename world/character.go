package world

import (
	. "github.com/marisvali/vlok/gamelib"
)

type Character struct {
	Pos       Pt
	Size      Pt
	MaxHealth Int
	Health    Int
}

func NewCharacter() (c Character) {
	c.MaxHealth = I(3)
	c.Health = c.MaxHealth
	return
}

func (c *Character) Step(w *World, input PlayerInput) {

}
