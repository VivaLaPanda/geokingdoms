package player

type Player struct {
	baseSpeed int
}

func NewPlayer() *Player {
	return &Player{
		baseSpeed: 1,
	}
}

func (p *Player) Speed() int {
	return p.baseSpeed
}
