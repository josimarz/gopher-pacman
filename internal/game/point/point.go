package point

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

func (p *Point) Equals(pt *Point) bool {
	return p.X == pt.X && p.Y == pt.Y
}

func (p *Point) Clone() *Point {
	return &Point{
		X: p.X,
		Y: p.Y,
	}
}
