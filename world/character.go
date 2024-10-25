package world

import (
	. "github.com/marisvali/vlok/gamelib"
)

type Character struct {
	Pos        Pt
	Size       Pt
	MaxHealth  Int
	Health     Int
	Picked     bool
	Speed      Int
	Ai         Ai
	MoveLimits Rectangle
}

func NewCharacter() (c Character) {
	c.MaxHealth = I(3)
	c.Health = c.MaxHealth
	c.Speed = U(5)
	c.MoveLimits = Rectangle{UPt(120, 90), UPt(790, 790)}
	return
}

func (c *Character) MoveToFood(w *World) {
	if c.Pos.DistTo(w.Food.Pos).Gt(U(3)) {
		dir := c.Pos.To(w.Food.Pos)
		dir.SetLen(c.Speed)
		c.Pos.Add(dir)
	}
}

func (c *Character) ChangePos(newPos Pt) {
	// Check if the new position is valid.
	if c.MoveLimits.ContainsPt(newPos) {
		c.Pos = newPos
	}
}

func (c *Character) Move(dir Pt) {
	dir.SetLen(c.Speed)
	c.ChangePos(c.Pos.Plus(dir))
}

func (c *Character) MoveLeft() {
	c.Move(UPt(-1, 0))
}

func (c *Character) MoveRight() {
	c.Move(UPt(1, 0))
}

func (c *Character) MoveUp() {
	c.Move(UPt(0, -1))
}

func (c *Character) MoveDown() {
	c.Move(UPt(0, 1))
}

func (c *Character) ClosestValidPos(pos Pt) Pt {
	x := Max(c.MoveLimits.Min().X, Min(pos.X, c.MoveLimits.Max().X))
	y := Max(c.MoveLimits.Min().Y, Min(pos.Y, c.MoveLimits.Max().Y))
	return Pt{x, y}
}

func (c *Character) Step(w *World, input PlayerInput) {
	if c.Picked {
		c.ChangePos(c.ClosestValidPos(input.Position))
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
