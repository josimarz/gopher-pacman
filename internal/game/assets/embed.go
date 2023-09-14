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

	//go:embed game_start.wav
	GameStart_wav []byte

	//go:embed munch_1.wav
	Munch1_wav []byte

	//go:embed munch_2.wav
	Munch2_wav []byte

	//go:embed power_pellet.wav
	PowerPellet_wav []byte

	//go:embed eat_ghost.wav
	EatGhost_wav []byte

	//go:embed death_1.wav
	Death1_wav []byte

	SpriteSheet *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(SpriteSheet_png))
	if err != nil {
		log.Fatal(err)
	}
	SpriteSheet = ebiten.NewImageFromImage(img)
}
