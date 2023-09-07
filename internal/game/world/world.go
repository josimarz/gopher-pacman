package world

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

const (
	Cols int = 21
	Rows int = 21
)

var (
	once       sync.Once
	world      *World
	contentSet = [][]tile.Content{
		{tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Pill, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Pill, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.None, tile.Wall, tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.None},
		{tile.None, tile.None, tile.None, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.None, tile.None, tile.None, tile.None, tile.None, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.None, tile.None, tile.None},
		{tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.Wall, tile.Wall, tile.Door, tile.Wall, tile.Wall, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall},
		{tile.None, tile.None, tile.None, tile.None, tile.None, tile.Dot, tile.None, tile.None, tile.Wall, tile.None, tile.None, tile.None, tile.Wall, tile.None, tile.None, tile.Dot, tile.None, tile.None, tile.None, tile.None, tile.None},
		{tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall},
		{tile.None, tile.None, tile.None, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.None, tile.None, tile.None, tile.None, tile.None, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.None, tile.None, tile.None},
		{tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Pill, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.None, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Pill, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.None},
		{tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.None},
	}
)

type World struct {
	tileSet []*tile.Tile
}

func Instance() *World {
	once.Do(func() {
		tileSet := []*tile.Tile{}
		for y := range contentSet {
			for x := range contentSet[y] {
				content := contentSet[y][x]
				point := point.New(x*tile.Size, y*tile.Size)
				tile := tile.New(content, point)
				tileSet = append(tileSet, tile)
			}
		}
		world = &World{
			tileSet: tileSet,
		}
	})
	return world
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, tile := range w.tileSet {
		tile.Draw(screen)
	}
}

func (w *World) Accessible(point *point.Point) bool {
	tile := w.TileAt(point)
	if tile == nil {
		return false
	}
	return tile.Accessible()
}

func (w *World) TileAt(point *point.Point) *tile.Tile {
	for _, tile := range w.tileSet {
		if tile.Point().Equals(point) {
			return tile
		}
	}
	return nil
}
