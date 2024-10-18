package world

import (
	"fmt"
	. "github.com/marisvali/vlok/gamelib"
	"math"
)

const Version = 1

type Food struct {
	Pos  Pt
	Size Pt
}

type World struct {
	Size      Pt
	Character Character
	Food      Food
	TimeStep  Int
}

type PlayerInput struct {
	Position Pt
	Pick     bool
	Release  bool
}

func NewWorld() (w World) {
	w.Character = NewCharacter()

	w.Size = UPt(800, 800)
	w.Character.Size = UPt(50, 50)
	w.Character.Pos = UPt(100, 200)
	w.Food.Size = UPt(50, 50)
	w.Food.Pos = UPt(400, 400)
	return
}

func (w *World) Step(input PlayerInput) {
	if input.Pick {
		if input.Position.DistTo(w.Character.Pos).Lt(U(50)) {
			w.Character.Pick()
		}
	}

	if input.Release {
		if w.Character.IsPicked() {
			w.Character.Release()
		}
	}

	w.Character.Step(w, input.Position)

	w.TimeStep.Inc()
	if w.TimeStep.Eq(I(math.MaxInt64)) {
		// Damn.
		Check(fmt.Errorf("got to an unusually large time step: %d", w.TimeStep.ToInt64()))
	}
}
