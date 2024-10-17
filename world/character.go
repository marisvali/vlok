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
	// Move towards the food.
	if c.Pos.DistTo(w.Food.Pos).Gt(U(3)) {
		dir := c.Pos.To(w.Food.Pos)
		dir.SetLen(U(1))
		c.Pos.Add(dir)
	}
}
