package point

import "github.com/josimarz/gopher-pacman/internal/game/direction"

type Point struct {
	X int
	Y int
}

func New(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Point) Clone() *Point {
	return &Point{
		X: p.X,
		Y: p.Y,
	}
}

func (p *Point) Dir(dst *Point) direction.Direction {
	if p.X < dst.X {
		return direction.Right
	}
	if p.X > dst.X {
		return direction.Left
	}
	if p.Y < dst.Y {
		return direction.Down
	}
	return direction.Up
}
