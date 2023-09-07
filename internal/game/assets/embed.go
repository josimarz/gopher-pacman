package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed spritesheet.png
	SpriteSheet_png []byte

	SpriteSheet *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(SpriteSheet_png))
	if err != nil {
		log.Fatal(err)
	}
	SpriteSheet = ebiten.NewImageFromImage(img)
}
