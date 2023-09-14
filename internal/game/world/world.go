package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

const (
	Cols int = 21
	Rows int = 21
)

var (
	world      *World
	contentSet = [][]tile.Content{
		{tile.Outline, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.PowerPellet, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.PowerPellet, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.None, tile.Wall, tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Outline},
		{tile.Outline, tile.Outline, tile.Outline, tile.Outline, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.None, tile.None, tile.None, tile.None, tile.None, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Outline, tile.Outline, tile.Outline, tile.Outline},
		{tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.Wall, tile.Wall, tile.Door, tile.Wall, tile.Wall, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall},
		{tile.None, tile.None, tile.None, tile.None, tile.None, tile.Dot, tile.None, tile.None, tile.Wall, tile.None, tile.None, tile.None, tile.Wall, tile.None, tile.None, tile.Dot, tile.None, tile.None, tile.None, tile.None, tile.None},
		{tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall},
		{tile.Outline, tile.Outline, tile.Outline, tile.Outline, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.None, tile.None, tile.None, tile.None, tile.None, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Outline, tile.Outline, tile.Outline, tile.Outline},
		{tile.Outline, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.None, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.None, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.PowerPellet, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.None, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.PowerPellet, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Dot, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Dot, tile.Wall, tile.Outline},
		{tile.Outline, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Wall, tile.Outline},
	}
)

func init() {
	tiles := []*tile.Tile{}
	for y := range contentSet {
		for x := range contentSet[y] {
			content := contentSet[y][x]
			point := point.New(x*tile.Size, y*tile.Size)
			tile := tile.New(content, point)
			tiles = append(tiles, tile)
		}
	}
	world = &World{
		tiles: tiles,
	}
}

func Draw(screen *ebiten.Image) {
	world.draw(screen)
}

func Reachable(point *point.Point) bool {
	return world.reachable(point)
}

func TileAt(point *point.Point) *tile.Tile {
	return world.tileAt(point)
}

type World struct {
	tiles []*tile.Tile
}

func (w *World) draw(screen *ebiten.Image) {
	for _, tile := range w.tiles {
		tile.Draw(screen)
	}
}

func (w *World) reachable(point *point.Point) bool {
	tile := w.tileAt(point)
	if tile == nil {
		return false
	}
	return tile.Reachable()
}

func (w *World) tileAt(point *point.Point) *tile.Tile {
	for _, tile := range w.tiles {
		if tile.Point().Equals(point) {
			return tile
		}
	}
	return nil
}
