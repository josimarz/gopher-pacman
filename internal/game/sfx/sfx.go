package sfx

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
)

const (
	sampleRate = 22050
)

var (
	audioCtx = audio.NewContext(sampleRate)
)

func PlayGameStart() {
	audioCtx.NewPlayerFromBytes(assets.GameStart_wav).Play()
}

func PlayMunch1() {
	audioCtx.NewPlayerFromBytes(assets.Munch1_wav).Play()
}

func PlayMunch2() {
	audioCtx.NewPlayerFromBytes(assets.Munch2_wav).Play()
}

func PlayPowerPellet() {
	audioCtx.NewPlayerFromBytes(assets.PowerPellet_wav).Play()
}

func PlayEatGhost() {
	audioCtx.NewPlayerFromBytes(assets.EatGhost_wav).Play()
}

func PlayDeath() {
	audioCtx.NewPlayerFromBytes(assets.Death1_wav).Play()
}
