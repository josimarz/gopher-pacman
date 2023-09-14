package ia

import (
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

func getDirections(src *point.Point) []*point.Point {
	return []*point.Point{
		point.New(src.X, src.Y-tile.Size),
		point.New(src.X, src.Y+tile.Size),
		point.New(src.X-tile.Size, src.Y),
		point.New(src.X+tile.Size, src.Y),
	}
}
