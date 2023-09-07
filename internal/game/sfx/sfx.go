package sfx

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
)

const (
	sampleRate int = 22050
)

var (
	once     sync.Once
	instance *audioPlayer
)

type audioPlayer struct {
	ctx *audio.Context
}

func AudioPlayer() *audioPlayer {
	once.Do(func() {
		instance = &audioPlayer{
			ctx: audio.NewContext(sampleRate),
		}
	})
	return instance
}

func (p *audioPlayer) PlayGameStart() {
	p.ctx.NewPlayerFromBytes(assets.GameStart_wav).Play()
}

func (p *audioPlayer) PlayPowerPellet() {
	p.ctx.NewPlayerFromBytes(assets.PowerPellet_wav).Play()
}

func (p *audioPlayer) PlayMunch1() {
	p.ctx.NewPlayerFromBytes(assets.Munch1_wav).Play()
}

func (p *audioPlayer) PlayMunch2() {
	p.ctx.NewPlayerFromBytes(assets.Munch2_wav).Play()
}
