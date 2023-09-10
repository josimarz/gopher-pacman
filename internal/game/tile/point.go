package tile

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
