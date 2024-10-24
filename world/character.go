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
	Speed     Int
	Ai        Ai
}

func NewCharacter() (c Character) {
	c.MaxHealth = I(3)
	c.Health = c.MaxHealth
	c.Speed = U(5)
	return
}

func (c *Character) MoveToFood(w *World) {
	if c.Pos.DistTo(w.Food.Pos).Gt(U(3)) {
		dir := c.Pos.To(w.Food.Pos)
		dir.SetLen(c.Speed)
		c.Pos.Add(dir)
	}
}

func (c *Character) MoveLeft(w *World) {
	dir := UPt(-1, 0)
	dir.SetLen(c.Speed)
	c.Pos.Add(dir)
}

func (c *Character) MoveRight(w *World) {
	dir := UPt(1, 0)
	dir.SetLen(c.Speed)
	c.Pos.Add(dir)
}

func (c *Character) MoveUp(w *World) {
	dir := UPt(0, -1)
	dir.SetLen(c.Speed)
	c.Pos.Add(dir)
}

func (c *Character) MoveDown(w *World) {
	dir := UPt(0, 1)
	dir.SetLen(c.Speed)
	c.Pos.Add(dir)
}

func (c *Character) Step(w *World, input PlayerInput) {
	if c.Picked {
		c.Pos = input.Position
	} else {
		c.Ai.Step(w, c, input)
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
