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
	Move    bool
	MovePt  Pt // tile-coordinates
	Shoot   bool
	ShootPt Pt // tile-coordinates
}

func NewWorld() (w World) {
	w.Character = NewCharacter()

	w.Size = Pt{I(800), I(800)}
	w.Character.Size = Pt{I(50), I(50)}
	w.Character.Pos = Pt{I(100), I(200)}
	w.Food.Size = Pt{I(50), I(50)}
	w.Food.Pos = Pt{I(400), I(400)}
	return
}

func (w *World) Step(input PlayerInput) {
	w.Character.Step(w, input)
	w.TimeStep.Inc()
	if w.TimeStep.Eq(I(math.MaxInt64)) {
		// Damn.
		Check(fmt.Errorf("got to an unusually large time step: %d", w.TimeStep.ToInt64()))
	}
}
