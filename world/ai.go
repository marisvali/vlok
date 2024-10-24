package world

type Ai struct {
	State AiState
}

type AiState int

const (
	MoveToFood AiState = iota
	MoveLeft
	MoveRight
	MoveUp
	MoveDown
)

func (a *Ai) Step(w *World, c *Character, input PlayerInput) {
	if input.MoveLeft {
		a.State = MoveLeft
	}

	if input.MoveRight {
		a.State = MoveRight
	}

	if input.MoveUp {
		a.State = MoveUp
	}

	if input.MoveDown {
		a.State = MoveDown
	}

	if input.MoveToFood {
		a.State = MoveToFood
	}

	switch a.State {
	case MoveToFood:
		c.MoveToFood(w)
	case MoveLeft:
		c.MoveLeft(w)
	case MoveRight:
		c.MoveRight(w)
	case MoveUp:
		c.MoveUp(w)
	case MoveDown:
		c.MoveDown(w)
	}
}
