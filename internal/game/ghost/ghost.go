package ghost

import (
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
	"github.com/josimarz/gopher-pacman/internal/game/direction"
	"github.com/josimarz/gopher-pacman/internal/game/ia"
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

type Color uint8

const (
	Red Color = iota
	Pink
	Cyan
	Orange
)

type Status uint8

const (
	Alive Status = iota
	Dizzy
	Recovering
	Dead
)

var (
	Blinky, Pinky, Inky, Clyde *Ghost
)

func init() {
	Blinky = new(Red)
	Pinky = new(Pink)
	Inky = new(Cyan)
	Clyde = new(Orange)
}

func Draw(screen *ebiten.Image) {
	Blinky.draw(screen)
	Pinky.draw(screen)
	Inky.draw(screen)
	Clyde.draw(screen)
}

func Update() {
	Blinky.move()
	Pinky.move()
	Inky.move()
	Clyde.move()
}

func DizzyAll() {
	go Blinky.dizzyMe()
	go Pinky.dizzyMe()
	go Inky.dizzyMe()
	go Clyde.dizzyMe()
}

type Ghost struct {
	goingHome bool
	path      ia.Stack[point.Point]
	status    Status
	color     Color
	dir       direction.Direction
	currPoint *point.Point
	nextPoint *point.Point
	speed     int
	dizziness []time.Time
}

func new(color Color) *Ghost {
	p := startPoint(color)
	return &Ghost{
		status:    Alive,
		color:     color,
		dir:       direction.Up,
		currPoint: p.Clone(),
		nextPoint: p.Clone(),
		speed:     1,
	}
}

func startPoint(color Color) *point.Point {
	switch color {
	case Red:
		return point.New(10*tile.Size, 7*tile.Size)
	case Pink:
		return point.New(10*tile.Size, 9*tile.Size)
	case Cyan:
		return point.New(9*tile.Size, 9*tile.Size)
	case Orange:
		return point.New(11*tile.Size, 9*tile.Size)
	default:
		log.Fatalf("invalid color: %v", color)
		return nil
	}
}

func (g *Ghost) CurrPoint() *point.Point {
	return g.currPoint
}

func (g *Ghost) Status() Status {
	return g.status
}

func (g *Ghost) Die() {
	g.status = Dead
}

func (g *Ghost) move() {
	if g.currPoint.Equals(g.nextPoint) {
		if g.status == Dead && !g.goingHome {
			g.goHome()
		}
		if g.path.Empty() {
			if g.goingHome {
				g.goingHome = false
			}
			if g.status == Dead {
				g.status = Alive
			}
			g.recreatePath()
		}
		g.nextPoint = g.path.Pop()
		g.dir = g.currPoint.Dir(g.nextPoint)
	}
	g.moveX()
	g.moveY()
}

func (g *Ghost) goHome() {
	g.goingHome = true
	goal := point.New(10*tile.Size, 9*tile.Size)
	g.path = ia.DFS(g.currPoint, goal)
}

func (g *Ghost) recreatePath() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	var goal *point.Point
	for {
		x := r.Intn(21)
		y := r.Intn(21)
		goal = point.New(x*tile.Size, y*tile.Size)
		if goal.Equals(g.currPoint) {
			continue
		}
		if world.Reachable(goal) {
			break
		}
	}
	g.path = ia.DFS(g.currPoint, goal)
}

func (g *Ghost) moveX() {
	if g.currPoint.X < g.nextPoint.X {
		g.currPoint.X += g.speed
	} else if g.currPoint.X > g.nextPoint.X {
		g.currPoint.X -= g.speed
	}
}

func (g *Ghost) moveY() {
	if g.currPoint.Y < g.nextPoint.Y {
		g.currPoint.Y += g.speed
	} else if g.currPoint.Y > g.nextPoint.Y {
		g.currPoint.Y -= g.speed
	}
}

func (g *Ghost) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.currPoint.X), float64(g.currPoint.Y))
	x, y := g.spritesCoords()
	sx, sy := x*tile.Size, y*tile.Size
	r := image.Rect(sx, sy, sx+tile.Size, sy+tile.Size)
	img := assets.SpriteSheet.SubImage(r).(*ebiten.Image)
	screen.DrawImage(img, op)
}

func (g *Ghost) spritesCoords() (int, int) {
	switch g.status {
	case Alive:
		return int(g.color), int(g.dir) + 2
	case Dizzy:
		return 5, 2
	case Recovering:
		return 5, 3
	case Dead:
		return 4, int(g.dir) + 2
	default:
		log.Fatalf("invalid status: %d", g.status)
		return 0, 0
	}
}

func (g *Ghost) dizzyMe() {
	if g.status == Dead {
		return
	}
	g.status = Dizzy
	ts := time.Now()
	g.dizziness = append(g.dizziness, ts)
	for _, t := range g.dizziness {
		if t.After(ts) {
			return
		}
	}
	time.Sleep(5 * time.Second)
	if g.status == Dead {
		return
	}
	g.status = Recovering
	time.Sleep(3 * time.Second)
	for _, t := range g.dizziness {
		if t.After(ts) {
			return
		}
	}
	if g.status == Dead {
		return
	}
	g.status = Alive
	g.dizziness = g.dizziness[:0]
}
