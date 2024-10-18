package world

import (
	. "github.com/marisvali/vlok/gamelib"
)

type Character struct {
	Pos       Pt
	Size      Pt
	MaxHealth Int
	Health    Int
	Picked    bool
}

func NewCharacter() (c Character) {
	c.MaxHealth = I(3)
	c.Health = c.MaxHealth
	return
}

func (c *Character) Step(w *World, playerPos Pt) {
	if c.Picked {
		c.Pos = playerPos
	} else {
		// Move towards the food.
		if c.Pos.DistTo(w.Food.Pos).Gt(U(3)) {
			dir := c.Pos.To(w.Food.Pos)
			dir.SetLen(U(1))
			c.Pos.Add(dir)
		}
	}
}

func (c *Character) Pick() {
	c.Picked = true
}

func (c *Character) Release() {
	c.Picked = false
}

func (c *Character) IsPicked() bool {
	return c.Picked
}
