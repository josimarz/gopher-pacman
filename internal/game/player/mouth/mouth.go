package mouth

type Status uint8

const (
	Open Status = iota + 1
	Closing
	Closed
	Opening
)

type Mouth struct {
	status Status
	delta  float64
	speed  float64
}

func New() *Mouth {
	return &Mouth{
		status: Closed,
		delta:  1,
		speed:  0.05,
	}
}

func (m *Mouth) Status() Status {
	return m.status
}

func (m *Mouth) Update() {
	m.delta += m.speed * m.delta
	if m.delta >= 5 {
		m.delta = 1
	}
	m.status = Status(m.delta)
}
