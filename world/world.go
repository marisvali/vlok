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
	Position   Pt
	Pick       bool
	Release    bool
	MoveLeft   bool
	MoveRight  bool
	MoveUp     bool
	MoveDown   bool
	MoveToFood bool
}

func NewWorld() (w World) {
	w.Character = NewCharacter()

	w.Size = UPt(900, 900)
	sz := 200
	w.Character.Size = UPt(sz, sz)
	w.Character.Pos = UPt(100, 200)
	w.Food.Size = UPt(200, 200)
	w.Food.Pos = UPt(450, 450)
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

	if input.MoveLeft {
		w.Character.State = MoveLeft
	}

	if input.MoveRight {
		w.Character.State = MoveRight
	}

	if input.MoveUp {
		w.Character.State = MoveUp
	}

	if input.MoveDown {
		w.Character.State = MoveDown
	}

	if input.MoveToFood {
		w.Character.State = MoveToFood
	}

	w.Character.Step(w, input.Position)

	w.TimeStep.Inc()
	if w.TimeStep.Eq(I(math.MaxInt64)) {
		// Damn.
		Check(fmt.Errorf("got to an unusually large time step: %d", w.TimeStep.ToInt64()))
	}
}
