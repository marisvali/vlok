package world

import (
	"fmt"
	. "github.com/marisvali/vlok/gamelib"
	"math"
)

const Version = 1

type World struct {
	Character Character
	TimeStep  Int
	Size      Pt
}

type PlayerInput struct {
	Move    bool
	MovePt  Pt // tile-coordinates
	Shoot   bool
	ShootPt Pt // tile-coordinates
}

func NewWorld() (w World) {
	w.Size = Pt{I(800), I(800)}
	w.Character = NewCharacter()
	w.Character.Pos = Pt{I(300), I(300)}
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
