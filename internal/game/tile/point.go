package tile

import (
	"github.com/josimarz/gopher-pacman/internal/game/direction"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) Equals(pt *Point) bool {
	return p.X == pt.X && p.Y == pt.Y
}

func (p *Point) Clone() *Point {
	return &Point{
		X: p.X,
		Y: p.Y,
	}
}

func (p *Point) Collide(pt *Point) bool {
	return (p.X <= pt.X && p.X+Size >= pt.X && p.Y <= pt.Y && p.Y+Size >= pt.Y) ||
		(pt.X <= p.X && pt.X+Size >= p.X && pt.Y <= p.Y && pt.Y+Size >= p.Y)
}

func (p *Point) Up() *Point {
	return &Point{
		X: p.X,
		Y: p.Y - Size,
	}
}

func (p *Point) Down() *Point {
	return &Point{
		X: p.X,
		Y: p.Y + Size,
	}
}

func (p *Point) Left() *Point {
	return &Point{
		X: p.X - Size,
		Y: p.Y,
	}
}

func (p *Point) Right() *Point {
	return &Point{
		X: p.X + Size,
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
