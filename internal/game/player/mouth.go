package player

type mouthStatus uint8

const (
	mouthOpen mouthStatus = iota + 1
	mouthClosing
	mouthClosed
	mouthOpening
)

func NewMouth() *Mouth {
	return &Mouth{
		status: mouthClosed,
		delta:  1,
		speed:  0.05,
	}
}

func (m *Mouth) update() {
	m.delta += m.speed * m.delta
	if m.delta >= 5 {
		m.delta = 1
	}
	m.status = mouthStatus(m.delta)
}
