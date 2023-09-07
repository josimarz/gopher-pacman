package position

type Position struct {
	X int
	Y int
}

func New(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func (p *Position) Equals(pos *Position) bool {
	return p.X == pos.X && p.Y == pos.Y
}
