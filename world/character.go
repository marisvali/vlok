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
	State     CharacterState
	Speed     Int
}

func NewCharacter() (c Character) {
	c.MaxHealth = I(3)
	c.Health = c.MaxHealth
	c.Speed = U(5)
	return
}

type CharacterState int

const (
	MoveToFood CharacterState = iota
	MoveLeft
	MoveRight
	MoveUp
	MoveDown
)

func (c *Character) MoveToFood(w *World, playerPos Pt) {
	if c.Pos.DistTo(w.Food.Pos).Gt(U(3)) {
		dir := c.Pos.To(w.Food.Pos)
		dir.SetLen(c.Speed)
		c.Pos.Add(dir)
	}
}

func (c *Character) MoveLeft(w *World, playerPos Pt) {
	dir := UPt(-1, 0)
	dir.SetLen(c.Speed)
	c.Pos.Add(dir)
}

func (c *Character) MoveRight(w *World, playerPos Pt) {
	dir := UPt(1, 0)
	dir.SetLen(c.Speed)
	c.Pos.Add(dir)
}

func (c *Character) MoveUp(w *World, playerPos Pt) {
	dir := UPt(0, -1)
	dir.SetLen(c.Speed)
	c.Pos.Add(dir)
}

func (c *Character) MoveDown(w *World, playerPos Pt) {
	dir := UPt(0, 1)
	dir.SetLen(c.Speed)
	c.Pos.Add(dir)
}

func (c *Character) Move(w *World, playerPos Pt) {
	switch c.State {
	case MoveToFood:
		c.MoveToFood(w, playerPos)
	case MoveLeft:
		c.MoveLeft(w, playerPos)
	case MoveRight:
		c.MoveRight(w, playerPos)
	case MoveUp:
		c.MoveUp(w, playerPos)
	case MoveDown:
		c.MoveDown(w, playerPos)
	}
}

func (c *Character) Step(w *World, playerPos Pt) {
	if c.Picked {
		c.Pos = playerPos
	} else {
		c.Move(w, playerPos)
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
